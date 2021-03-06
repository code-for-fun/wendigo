package btree

//	A cursor is a pointer to a particular entry within a particular b-tree within a database file.
//	The entry is identified by its MemoryPage and the index in MemoryPage.CellIndices[] of the entry.
//	A single database file can be shared by two more database connections, but cursors cannot be shared.  Each cursor is associated with a
//	particular database connection identified Cursor.pBtree.db.
//	Fields in this structure are accessed under the BtShared.mutex found at self.pBt.mutex.
struct Cursor {
	pBtreee				*Btree			//	The Btree to which this cursor belongs
	pBt					*BtShared		//	The BtShared this cursor points to
	Next, pPrev		*Cursor			//	Forms a linked list of all cursors
	pKeyInfo			*KeyInfo		//	Argument passed to comparison function
	OverflowCache		*PageNumber		//	Cache of overflow page locations
	RootPage			PageNumber		//	The root page of this tree
	cachedRowid			int64	//	Next rowid cache.  0 means not valid
	CellInfo							//	A parse of the cell we are pointing at
	nKey				int64			//	Size of pKey, or last integer key
	pKey				*void			//	Saved key that was cursor's last known position
	skiNext			int				//	Prev() is noop if negative. Next() is noop if positive
	Writable			bool			//	True if writable
	atLast				byte			//	Cursor pointing to the last entry
	validNKey			byte			//	True if info.nKey is valid
	eState				byte			//	One of the CURSOR_XXX constants (see below)
	isIncrblobHandle	byte			//	True if this cursor is an incr. io handle */
	iPage				int16			//	Index of current page in apPage
	aiIdx				[]uint16		//	Current index in apPage[i]
	Pages				[]*MemoryPage	//	Pages from root to current page
}

//	Potential values for Cursor.eState.
//
//	CURSOR_VALID:
//		Cursor points to a valid entry. getPayload() etc. may be called.
//
//	CURSOR_INVALID:
//		Cursor does not point to a valid entry. This can happen (for example) because the table is empty or because BtreeCursorFirst() has not been called.
//
//	CURSOR_REQUIRESEEK:
//		The table that this cursor was opened on still exists, but has been modified since the cursor was last used. The cursor position is saved
//		in variables Cursor.pKey and Cursor.nKey. When a cursor is in this state, restoreCursorPosition() can be called to attempt to
//		seek the cursor to the saved position.
//
//	CURSOR_FAULT:
//		A unrecoverable error (an I/O error or a malloc failure) has occurred on a different connection that shares the BtShared cache with this
//		cursor. The error has left the cache in an inconsistent state. Do nothing else with this cursor. Any attempt to use the cursor
//		should return the error code stored in Cursor.skip
const(
	CURSOR_INVALID		= iota
	CURSOR_VALID
	CURSOR_REQUIRESEEK
	CURSOR_FAULT
)

//	The page that pCur currently points to has just been modified in some way. This function figures out if this modification means the
//	tree needs to be balanced, and if so calls the appropriate balancing routine. Balancing routines are:
//
//		balance_quick()
//		balance_deeper()
//		BalanceNonroot()

func (pCur *Cursor) Balance() (rc int) {
	aBalanceQuickSpace	[13]byte
	pFree				*byte

	rc = SQLITE_OK
	nMin := pCur.pBt.usableSize * 2 / 3

	for {
		iPage := pCur.iPage
 		pPage := pCur.Pages[iPage]

		switch {
		case iPage == 0:
			if pPage.nOverflow {
				//	The root page of the b-tree is overfull. In this case call the balance_deeper() function to create a new child for the root-page
				//	and copy the current contents of the root-page to it. The next iteration of the do-loop will balance the child page.
				assert( (balance_deeper_called++)==0 );
				rc = balance_deeper(pPage, &pCur.Pages[1]);
				if rc == SQLITE_OK {
					pCur.iPage = 1
					pCur.aiIdx[0] = 0
					pCur.aiIdx[1] = 0
					assert( pCur.Pages[1].nOverflow )
				}
			} else {
				break
			}
		case pPage.nOverflow == 0 && pPage.nFree <= nMin:
			break;
		default:
			MemoryPage * const pParent = pCur.Pages[iPage-1]
			int const iIdx = pCur.aiIdx[iPage-1]

			rc = pParent.DbPage.Write()
			if rc == SQLITE_OK {
				if pPage.HasData && pPage.nOverflow == 1 && pPage.aiOvfl[0] == pPage.nCell && pParent.pgno != 1 && pParent.nCell == iIdx) {
					//	Call balance_quick() to create a new sibling of pPage on which to store the overflow cell. balance_quick() inserts a new cell
					//	into pParent, which may cause pParent overflow. If this happens, the next interation of the do-loop will balance pParent
					//	use either BalanceNonroot() or balance_deeper(). Until this happens, the overflow cell is stored in the aBalanceQuickSpace[] buffer.
					//
					//	The purpose of the following assert() is to check that only a single call to balance_quick() is made for each call to this
					//	function. If this were not verified, a subtle bug involving reuse of the aBalanceQuickSpace[] might sneak in.
					assert( (balance_quick_called++)==0 )
					rc = balance_quick(pParent, pPage, aBalanceQuickSpace)
				} else {
					//	In this case, call BalanceNonroot() to redistribute cells between pPage and up to 2 of its sibling pages. This involves
					//	modifying the contents of pParent, which may cause pParent to become overfull or underfull. The next iteration of the do-loop
					//	will balance the parent page to correct this.
					//
					//	If the parent page becomes overfull, the overflow cell or cells are stored in the pSpace buffer allocated immediately below.
					//	A subsequent iteration of the do-loop will deal with this by calling BalanceNonroot() (balance_deeper() may be called first,
					//	but it doesn't deal with overflow cells - just moves them to a different page). Once this subsequent call to BalanceNonroot()
					//	has completed, it is safe to release the pSpace buffer used by the previous call, as the overflow cell data will have been
					//	copied either into the body of a database page or into the new pSpace buffer passed to the latter call to BalanceNonroot().
					byte *pSpace = sqlite3PageMalloc(pCur.pBt.pageSize)
					rc = pParent.BalanceNonroot(iIdx, pSpace, iPage == 1)
					if pFree {
						//	If pFree is not NULL, it points to the pSpace buffer used by a previous call to BalanceNonroot(). Its contents are
						//	now stored either on real database pages or within the new pSpace buffer, so it may be safely freed here.
						sqlite3PageFree(pFree);
					}

					//	The pSpace buffer will be freed after the next call to BalanceNonroot(), or just before this function returns, whichever comes first.
					pFree = pSpace
				}
			}
			pPage.nOverflow = 0

			//	The next iteration of the do-loop balances the parent page.
			pPage.Release()
			pCur.iPage--
		}
		if rc == SQLITE_OK {
			break
		}
	}

	if pFree {
		sqlite3PageFree(pFree);
	}
	return rc
}

//	Move the cursor down to a new child page. The newPageNumber argument is the page number of the child page to move to.
//
//	This function returns SQLITE_CORRUPT if the page-header flags field of the new child page does not match the flags field of the parent (i.e.
//	if an IsIntegerKey page appears to be the parent of a non-IsIntegerKey page, or vice-versa).

func (pCur *Cursor) MoveToChild(newPageNumber uint32) (rc int) {
	pNewPage	*MemoryPage

	i := pCur.iPage
	pBt := pCur.pBt

	assert( pCur.eState == CURSOR_VALID )
	assert( pCur.iPage < BTCURSOR_MAX_DEPTH )
	if pCur.iPage >= (BTCURSOR_MAX_DEPTH - 1) {
		rc = SQLITE_CORRUPT_BKPT
	} else {
		if rc, pNewPage = pBt.GetPageAndInitialize(newPageNumber); rc == SQLITE_OK {
			pCur.Pages[i+1] = pNewPage
			pCur.aiIdx[i+1] = 0
			pCur.iPage++

			pCur.CellInfo.nSize = 0
			pCur.validNKey = 0
			if pNewPage.nCell < 1 || pNewPage.IsIntegerKey != pCur.Pages[i].IsIntegerKey {
				rc = SQLITE_CORRUPT_BKPT
			}
		}
	}
	return
}


//	Delete the entry that the cursor is pointing to. The cursor is left pointing at a arbitrary location.
func (pCur *Cursor) Delete() (rc int) {
int sqlite3BtreeDelete(Cursor *pCur){
	p := pCur.pBtree
	pBt = p.pBt

	MemoryPage *pPage;                      /* Page to delete cell from */
	unsigned char *pCell;                /* Pointer to cell to delete */
	int iCellIdx;                        /* Index of cell to delete */
	int iCellDepth;                      /* Depth of node containing pCell */

	assert( pBt.inTransaction == TRANS_WRITE )
	assert( pBt.Flags & BTS_READ_ONLY == 0 )
	assert( pCur.Writable )
	assert( hasSharedCacheTableLock(p, pCur.RootPage, pCur.pKeyInfo != 0, 2) )
	assert( !hasReadConflicts(p, pCur.RootPage) )

	if pCur.aiIdx[pCur.iPage] >= pCur.Pages[pCur.iPage].nCell || pCur.eState!=CURSOR_VALID {
		return SQLITE_ERROR;  /* Something has gone awry. */
	}

	iCellDepth = pCur.iPage
	iCellIdx = pCur.aiIdx[iCellDepth]
	pPage = pCur.Pages[iCellDepth]
	pCell = pPage.FindCell(iCellIdx)

	//	If the page containing the entry to delete is not a IsLeaf page, move the cursor to the largest entry in the tree that is smaller than the entry being deleted. This cell will replace the cell being deleted from the internal node. The 'previous' entry is used for this instead of the 'next' entry, as the previous entry is always a part of the sub-tree headed by the child page of the cell being deleted. This makes balancing the tree following the delete operation easier.
	if !pPage.IsLeaf {
		int notUsed;
		if rc = sqlite3BtreePrevious(pCur, &notUsed) {
			return
		}
	}

	//	Save the positions of any other cursors open on this table before making any modifications. Make the page containing the entry to be deleted writable. Then free any overflow pages associated with the entry and finally remove the cell itself from within the page.
	if rc = saveAllCursors(pBt, pCur.RootPage, pCur); rc != SQLITE_OK {
		return
	}

	//	If this is a delete operation to remove a row from a table b-tree, invalidate any incrblob cursors open on the row being deleted.
	if pCur.pKeyInfo == nil {
		p.InvalidateIncrblobCursors(pCur.CellInfo.nKey, false)
	}

	if rc = pPage.DbPage.Write(); rc != SQLITE_OK {
		return
	}
	rc = clearCell(pPage, pCell)
	if dropCell(pPage, iCellIdx, pPage.cellSize(pCell), &rc); rc != SQLITE_OK {
		return
	}

	//	If the cell deleted was not located on a IsLeaf page, then the cursor is currently pointing to the largest entry in the sub-tree headed by the child-page of the cell that was just deleted from an internal node. The cell from the IsLeaf node needs to be moved to the internal node to replace the deleted cell.
	if !pPage.IsLeaf {
		pIsLeaf := pCur.Pages[pCur.iPage]
		int nCell
		n := pCur.Pages[iCellDepth + 1].pgno
		unsigned char *pTmp

		pCell = pIsLeaf.FindCell(pIsLeaf.nCell - 1)
		nCell = pIsLeaf.cellSize(pCell)
		assert( MX_CELL_SIZE(pBt) >= nCell )

		allocateTempSpace(pBt)
		pTmp = pBt.pTmpSpace

		if rc = pIsLeaf.DbPage.Write(); rc == SQLITE_OK {
			//	TODO:	pCell needs to point to memory 4 bytes before its start in the page
			if rc = pPage.InsertCell(iCellIdx, pCell[-4:], pTmp, n); rc == SQLITE_OK {
				dropCell(pIsLeaf, pIsLeaf.nCell - 1, nCell, &rc)
			}
		}
		if rc != SQLITE_OK {
			return
		}
	}

	//	Balance the tree. If the entry deleted was located on a IsLeaf page, then the cursor still points to that page. In this case the first call to Balance() repairs the tree, and the if(...) condition is never true.
	//	Otherwise, if the entry deleted was on an internal node page, then pCur is pointing to the IsLeaf page from which a cell was removed to replace the cell deleted from the internal node. This is slightly tricky as the IsLeaf node may be underfull, and the internal node may be either under or overfull. In this case run the balancing algorithm on the IsLeaf node first. If the balance proceeds far enough up the tree that we can be sure that any problem in the internal node has been corrected, so be it. Otherwise, after balancing the IsLeaf node, walk the cursor up the tree to the internal node and balance it as well.
	if rc = pCur.Balance(); rc == SQLITE_OK && pCur.iPage > iCellDepth {
		for pCur.iPage > iCellDepth {
			pCur.iPage--
			pCur.Pages[pCur.iPage].Release()
		}
		rc = pCur.Balance()
	}
	if rc == SQLITE_OK {
		pCur.MoveToRoot()
	}
	return
}

//	Set a flag on this cursor to cache the locations of pages from the overflow list for the current row. This is used by cursors opened for incremental blob IO only.
//	This function sets a flag only. The actual page location cache (stored in Cursor.OverflowCache[]) is allocated and used by function accessPayload() (the worker function for sqlite3BtreeData() and sqlite3BtreePutData()).
func (c *Cursor) CacheOverflow() {
	c.CriticalSection(func() {
		c.OverflowCache = nil
		c.isIncrblobHandle = 1
	})
}
