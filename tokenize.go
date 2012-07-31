/* An tokenizer for SQL
**
** This file contains C code that splits an SQL input string up into
** individual tokens and sends those tokens one-by-one over to the
** parser for analysis.
*/
/* #include <stdlib.h> */

/*
** The charMap() macro maps alphabetic characters into their
** lower-case ASCII equivalent.
*/
# define charMap(X) sqlite3UpperToLower[(unsigned char)X]

/*
** The sqlite3KeywordCode function looks up an identifier to determine if
** it is a keyword.  If it is a keyword, the token code of that keyword is 
** returned.  If the input is not a keyword, TK_ID is returned.
**
** The implementation of this routine was generated by a program,
** mkkeywordhash.h, located in the tool subdirectory of the distribution.
** The output of the mkkeywordhash.c program is written into a file
** named keywordhash.h and then included into this source file by
** the #include below.
*/
/************** Include keywordhash.h in the middle of tokenize.c ************/
/************** Begin file keywordhash.h *************************************/
/***** This file contains automatically generated code ******
**
** The code in this file has been automatically generated by
**
**   sqlite/tool/mkkeywordhash.c
**
** The code in this file implements a function that determines whether
** or not a given identifier is really an SQL keyword.  The same thing
** might be implemented more directly using a hand-written hash table.
** But by using this automatically generated code, the size of the code
** is substantially reduced.  This is important for embedded applications
** on platforms with limited memory.
*/
/* Hash score: 175 */
static int keywordCode(const char *z, int n){
  /* zText[] encodes 811 bytes of keywords in 541 bytes */
  /*   REINDEXEDESCAPEACHECKEYBEFOREIGNOREGEXPLAINSTEADDATABASELECT       */
  /*   ABLEFTHENDEFERRABLELSEXCEPTRANSACTIONATURALTERAISEXCLUSIVE         */
  /*   XISTSAVEPOINTERSECTRIGGEREFERENCESCONSTRAINTOFFSETEMPORARY         */
  /*   UNIQUERYATTACHAVINGROUPDATEBEGINNERELEASEBETWEENOTNULLIKE          */
  /*   CASCADELETECASECOLLATECREATECURRENT_DATEDETACHIMMEDIATEJOIN        */
  /*   SERTMATCHPLANALYZEPRAGMABORTVALUESVIRTUALIMITWHENWHERENAME         */
  /*   AFTEREPLACEANDEFAULTAUTOINCREMENTCASTCOLUMNCOMMITCONFLICTCROSS     */
  /*   CURRENT_TIMESTAMPRIMARYDEFERREDISTINCTDROPFAILFROMFULLGLOBYIF      */
  /*   ISNULLORDERESTRICTOUTERIGHTROLLBACKROWUNIONUSINGVACUUMVIEW         */
  /*   INITIALLY                                                          */
  static const char zText[540] = {
    'R','E','I','N','D','E','X','E','D','E','S','C','A','P','E','A','C','H',
    'E','C','K','E','Y','B','E','F','O','R','E','I','G','N','O','R','E','G',
    'E','X','P','L','A','I','N','S','T','E','A','D','D','A','T','A','B','A',
    'S','E','L','E','C','T','A','B','L','E','F','T','H','E','N','D','E','F',
    'E','R','R','A','B','L','E','L','S','E','X','C','E','P','T','R','A','N',
    'S','A','C','T','I','O','N','A','T','U','R','A','L','T','E','R','A','I',
    'S','E','X','C','L','U','S','I','V','E','X','I','S','T','S','A','V','E',
    'P','O','I','N','T','E','R','S','E','C','T','R','I','G','G','E','R','E',
    'F','E','R','E','N','C','E','S','C','O','N','S','T','R','A','I','N','T',
    'O','F','F','S','E','T','E','M','P','O','R','A','R','Y','U','N','I','Q',
    'U','E','R','Y','A','T','T','A','C','H','A','V','I','N','G','R','O','U',
    'P','D','A','T','E','B','E','G','I','N','N','E','R','E','L','E','A','S',
    'E','B','E','T','W','E','E','N','O','T','N','U','L','L','I','K','E','C',
    'A','S','C','A','D','E','L','E','T','E','C','A','S','E','C','O','L','L',
    'A','T','E','C','R','E','A','T','E','C','U','R','R','E','N','T','_','D',
    'A','T','E','D','E','T','A','C','H','I','M','M','E','D','I','A','T','E',
    'J','O','I','N','S','E','R','T','M','A','T','C','H','P','L','A','N','A',
    'L','Y','Z','E','P','R','A','G','M','A','B','O','R','T','V','A','L','U',
    'E','S','V','I','R','T','U','A','L','I','M','I','T','W','H','E','N','W',
    'H','E','R','E','N','A','M','E','A','F','T','E','R','E','P','L','A','C',
    'E','A','N','D','E','F','A','U','L','T','A','U','T','O','I','N','C','R',
    'E','M','E','N','T','C','A','S','T','C','O','L','U','M','N','C','O','M',
    'M','I','T','C','O','N','F','L','I','C','T','C','R','O','S','S','C','U',
    'R','R','E','N','T','_','T','I','M','E','S','T','A','M','P','R','I','M',
    'A','R','Y','D','E','F','E','R','R','E','D','I','S','T','I','N','C','T',
    'D','R','O','P','F','A','I','L','F','R','O','M','F','U','L','L','G','L',
    'O','B','Y','I','F','I','S','N','U','L','L','O','R','D','E','R','E','S',
    'T','R','I','C','T','O','U','T','E','R','I','G','H','T','R','O','L','L',
    'B','A','C','K','R','O','W','U','N','I','O','N','U','S','I','N','G','V',
    'A','C','U','U','M','V','I','E','W','I','N','I','T','I','A','L','L','Y',
  };
  static const unsigned char aHash[127] = {
      72, 101, 114,  70,   0,  45,   0,   0,  78,   0,  73,   0,   0,
      42,  12,  74,  15,   0, 113,  81,  50, 108,   0,  19,   0,   0,
     118,   0, 116, 111,   0,  22,  89,   0,   9,   0,   0,  66,  67,
       0,  65,   6,   0,  48,  86,  98,   0, 115,  97,   0,   0,  44,
       0,  99,  24,   0,  17,   0, 119,  49,  23,   0,   5, 106,  25,
      92,   0,   0, 121, 102,  56, 120,  53,  28,  51,   0,  87,   0,
      96,  26,   0,  95,   0,   0,   0,  91,  88,  93,  84, 105,  14,
      39, 104,   0,  77,   0,  18,  85, 107,  32,   0, 117,  76, 109,
      58,  46,  80,   0,   0,  90,  40,   0, 112,   0,  36,   0,   0,
      29,   0,  82,  59,  60,   0,  20,  57,   0,  52,
  };
  static const unsigned char aNext[121] = {
       0,   0,   0,   0,   4,   0,   0,   0,   0,   0,   0,   0,   0,
       0,   2,   0,   0,   0,   0,   0,   0,  13,   0,   0,   0,   0,
       0,   7,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,   0,
       0,   0,   0,   0,  33,   0,  21,   0,   0,   0,  43,   3,  47,
       0,   0,   0,   0,  30,   0,  54,   0,  38,   0,   0,   0,   1,
      62,   0,   0,  63,   0,  41,   0,   0,   0,   0,   0,   0,   0,
      61,   0,   0,   0,   0,  31,  55,  16,  34,  10,   0,   0,   0,
       0,   0,   0,   0,  11,  68,  75,   0,   8,   0, 100,  94,   0,
     103,   0,  83,   0,  71,   0,   0, 110,  27,  37,  69,  79,   0,
      35,  64,   0,   0,
  };
  static const unsigned char aLen[121] = {
       7,   7,   5,   4,   6,   4,   5,   3,   6,   7,   3,   6,   6,
       7,   7,   3,   8,   2,   6,   5,   4,   4,   3,  10,   4,   6,
      11,   6,   2,   7,   5,   5,   9,   6,   9,   9,   7,  10,  10,
       4,   6,   2,   3,   9,   4,   2,   6,   5,   6,   6,   5,   6,
       5,   5,   7,   7,   7,   3,   2,   4,   4,   7,   3,   6,   4,
       7,   6,  12,   6,   9,   4,   6,   5,   4,   7,   6,   5,   6,
       7,   5,   4,   5,   6,   5,   7,   3,   7,  13,   2,   2,   4,
       6,   6,   8,   5,  17,  12,   7,   8,   8,   2,   4,   4,   4,
       4,   4,   2,   2,   6,   5,   8,   5,   5,   8,   3,   5,   5,
       6,   4,   9,   3,
  };
  static const unsigned short int aOffset[121] = {
       0,   2,   2,   8,   9,  14,  16,  20,  23,  25,  25,  29,  33,
      36,  41,  46,  48,  53,  54,  59,  62,  65,  67,  69,  78,  81,
      86,  91,  95,  96, 101, 105, 109, 117, 122, 128, 136, 142, 152,
     159, 162, 162, 165, 167, 167, 171, 176, 179, 184, 189, 194, 197,
     203, 206, 210, 217, 223, 223, 223, 226, 229, 233, 234, 238, 244,
     248, 255, 261, 273, 279, 288, 290, 296, 301, 303, 310, 315, 320,
     326, 332, 337, 341, 344, 350, 354, 361, 363, 370, 372, 374, 383,
     387, 393, 399, 407, 412, 412, 428, 435, 442, 443, 450, 454, 458,
     462, 466, 469, 471, 473, 479, 483, 491, 495, 500, 508, 511, 516,
     521, 527, 531, 536,
  };
  static const unsigned char aCode[121] = {
    TK_REINDEX,    TK_INDEXED,    TK_INDEX,      TK_DESC,       TK_ESCAPE,     
    TK_EACH,       TK_CHECK,      TK_KEY,        TK_BEFORE,     TK_FOREIGN,    
    TK_FOR,        TK_IGNORE,     TK_LIKE_KW,    TK_EXPLAIN,    TK_INSTEAD,    
    TK_ADD,        TK_DATABASE,   TK_AS,         TK_SELECT,     TK_TABLE,      
    TK_JOIN_KW,    TK_THEN,       TK_END,        TK_DEFERRABLE, TK_ELSE,       
    TK_EXCEPT,     TK_TRANSACTION,TK_ACTION,     TK_ON,         TK_JOIN_KW,    
    TK_ALTER,      TK_RAISE,      TK_EXCLUSIVE,  TK_EXISTS,     TK_SAVEPOINT,  
    TK_INTERSECT,  TK_TRIGGER,    TK_REFERENCES, TK_CONSTRAINT, TK_INTO,       
    TK_OFFSET,     TK_OF,         TK_SET,        TK_TEMP,       TK_TEMP,       
    TK_OR,         TK_UNIQUE,     TK_QUERY,      TK_ATTACH,     TK_HAVING,     
    TK_GROUP,      TK_UPDATE,     TK_BEGIN,      TK_JOIN_KW,    TK_RELEASE,    
    TK_BETWEEN,    TK_NOTNULL,    TK_NOT,        TK_NO,         TK_NULL,       
    TK_LIKE_KW,    TK_CASCADE,    TK_ASC,        TK_DELETE,     TK_CASE,       
    TK_COLLATE,    TK_CREATE,     TK_CTIME_KW,   TK_DETACH,     TK_IMMEDIATE,  
    TK_JOIN,       TK_INSERT,     TK_MATCH,      TK_PLAN,       TK_ANALYZE,    
    TK_PRAGMA,     TK_ABORT,      TK_VALUES,     TK_VIRTUAL,    TK_LIMIT,      
    TK_WHEN,       TK_WHERE,      TK_RENAME,     TK_AFTER,      TK_REPLACE,    
    TK_AND,        TK_DEFAULT,    TK_AUTOINCR,   TK_TO,         TK_IN,         
    TK_CAST,       TK_COLUMNKW,   TK_COMMIT,     TK_CONFLICT,   TK_JOIN_KW,    
    TK_CTIME_KW,   TK_CTIME_KW,   TK_PRIMARY,    TK_DEFERRED,   TK_DISTINCT,   
    TK_IS,         TK_DROP,       TK_FAIL,       TK_FROM,       TK_JOIN_KW,    
    TK_LIKE_KW,    TK_BY,         TK_IF,         TK_ISNULL,     TK_ORDER,      
    TK_RESTRICT,   TK_JOIN_KW,    TK_JOIN_KW,    TK_ROLLBACK,   TK_ROW,        
    TK_UNION,      TK_USING,      TK_VACUUM,     TK_VIEW,       TK_INITIALLY,  
    TK_ALL,        
  };
  int h, i;
  if( n<2 ) return TK_ID;
  h = ((charMap(z[0])*4) ^
      (charMap(z[n-1])*3) ^
      n) % 127;
  for(i=((int)aHash[h])-1; i>=0; i=((int)aNext[i])-1){
    if aLen[i] == n && CaseInsensitiveMatchN(&zText[aOffset[i]],z,n) {
      return aCode[i];
    }
  }
  return TK_ID;
}
 int sqlite3KeywordCode(const unsigned char *z, int n){
  return keywordCode((char*)z, n);
}
#define SQLITE_N_KEYWORD 121

/************** End of keywordhash.h *****************************************/
/************** Continuing where we left off in tokenize.c *******************/


/*
** If X is a character that can be used in an identifier then
** IdChar(X) will be true.  Otherwise it is false.
**
** For ASCII, any character with the high-order bit set is
** allowed in an identifier.  For 7-bit characters, 
** sqlite3IsIdChar[X] must be 1.
**
** Ticket #1066.  the SQL standard does not allow '$' in the
** middle of identfiers.  But many SQL implementations do. 
** SQLite will allow '$' in identifiers for compatibility.
** But the feature is undocumented.
*/
#define IdChar(C)  ((sqlite3CtypeMap[(unsigned char)C]&0x46)!=0)

/*
** Return the length of the token that begins at z[0]. 
** Store the token type in *tokenType before returning.
*/
 int sqlite3GetToken(const unsigned char *z, int *tokenType){
  int i, c;
  switch( *z ){
    case ' ': case '\t': case '\n': case '\f': case '\r': {
      for(i=1; sqlite3Isspace(z[i]); i++){}
      *tokenType = TK_SPACE;
      return i;
    }
    case '-': {
      if( z[1]=='-' ){
        /* IMP: R-50417-27976 -- syntax diagram for comments */
        for(i=2; (c=z[i])!=0 && c!='\n'; i++){}
        *tokenType = TK_SPACE;   /* IMP: R-22934-25134 */
        return i;
      }
      *tokenType = TK_MINUS;
      return 1;
    }
    case '(': {
      *tokenType = TK_LP;
      return 1;
    }
    case ')': {
      *tokenType = TK_RP;
      return 1;
    }
    case ';': {
      *tokenType = TK_SEMI;
      return 1;
    }
    case '+': {
      *tokenType = TK_PLUS;
      return 1;
    }
    case '*': {
      *tokenType = TK_STAR;
      return 1;
    }
    case '/': {
      if( z[1]!='*' || z[2]==0 ){
        *tokenType = TK_SLASH;
        return 1;
      }
      /* IMP: R-50417-27976 -- syntax diagram for comments */
      for(i=3, c=z[2]; (c!='*' || z[i]!='/') && (c=z[i])!=0; i++){}
      if( c ) i++;
      *tokenType = TK_SPACE;   /* IMP: R-22934-25134 */
      return i;
    }
    case '%': {
      *tokenType = TK_REM;
      return 1;
    }
    case '=': {
      *tokenType = TK_EQ;
      return 1 + (z[1]=='=');
    }
    case '<': {
      if( (c=z[1])=='=' ){
        *tokenType = TK_LE;
        return 2;
      }else if( c=='>' ){
        *tokenType = TK_NE;
        return 2;
      }else if( c=='<' ){
        *tokenType = TK_LSHIFT;
        return 2;
      }else{
        *tokenType = TK_LT;
        return 1;
      }
    }
    case '>': {
      if( (c=z[1])=='=' ){
        *tokenType = TK_GE;
        return 2;
      }else if( c=='>' ){
        *tokenType = TK_RSHIFT;
        return 2;
      }else{
        *tokenType = TK_GT;
        return 1;
      }
    }
    case '!': {
      if( z[1]!='=' ){
        *tokenType = TK_ILLEGAL;
        return 2;
      }else{
        *tokenType = TK_NE;
        return 2;
      }
    }
    case '|': {
      if( z[1]!='|' ){
        *tokenType = TK_BITOR;
        return 1;
      }else{
        *tokenType = TK_CONCAT;
        return 2;
      }
    }
    case ',': {
      *tokenType = TK_COMMA;
      return 1;
    }
    case '&': {
      *tokenType = TK_BITAND;
      return 1;
    }
    case '~': {
      *tokenType = TK_BITNOT;
      return 1;
    }
    case '`':
    case '\'':
    case '"': {
      int delim = z[0];
      for(i=1; (c=z[i])!=0; i++){
        if( c==delim ){
          if( z[i+1]==delim ){
            i++;
          }else{
            break;
          }
        }
      }
      if( c=='\'' ){
        *tokenType = TK_STRING;
        return i+1;
      }else if( c!=0 ){
        *tokenType = TK_ID;
        return i+1;
      }else{
        *tokenType = TK_ILLEGAL;
        return i;
      }
    }
    case '.': {
      if( !sqlite3Isdigit(z[1]) )
      {
        *tokenType = TK_DOT;
        return 1;
      }
      /* If the next character is a digit, this is a floating point
      ** number that begins with ".".  Fall thru into the next case */
    }
    case '0': case '1': case '2': case '3': case '4':
    case '5': case '6': case '7': case '8': case '9': {
      *tokenType = TK_INTEGER;
      for(i=0; sqlite3Isdigit(z[i]); i++){}
      if( z[i]=='.' ){
        i++;
        while( sqlite3Isdigit(z[i]) ){ i++; }
        *tokenType = TK_FLOAT;
      }
      if( (z[i]=='e' || z[i]=='E') &&
           ( sqlite3Isdigit(z[i+1]) 
            || ((z[i+1]=='+' || z[i+1]=='-') && sqlite3Isdigit(z[i+2]))
           )
      ){
        i += 2;
        while( sqlite3Isdigit(z[i]) ){ i++; }
        *tokenType = TK_FLOAT;
      }
      while( IdChar(z[i]) ){
        *tokenType = TK_ILLEGAL;
        i++;
      }
      return i;
    }
    case '[': {
      for(i=1, c=z[0]; c!=']' && (c=z[i])!=0; i++){}
      *tokenType = c==']' ? TK_ID : TK_ILLEGAL;
      return i;
    }
    case '?': {
      *tokenType = TK_VARIABLE;
      for(i=1; sqlite3Isdigit(z[i]); i++){}
      return i;
    }
    case '#': {
      for(i=1; sqlite3Isdigit(z[i]); i++){}
      if( i>1 ){
        /* Parameters of the form #NNN (where NNN is a number) are used
        ** internally by sqlite3NestedParse.  */
        *tokenType = TK_REGISTER;
        return i;
      }
      /* Fall through into the next case if the '#' is not followed by
      ** a digit. Try to match #AAAA where AAAA is a parameter name. */
    }
#ifndef SQLITE_OMIT_TCL_VARIABLE
    case '$':
#endif
    case '@':  /* For compatibility with MS SQL Server */
    case ':': {
      int n = 0;
      *tokenType = TK_VARIABLE;
      for(i=1; (c=z[i])!=0; i++){
        if( IdChar(c) ){
          n++;
#ifndef SQLITE_OMIT_TCL_VARIABLE
        }else if( c=='(' && n>0 ){
          do{
            i++;
          }while( (c=z[i])!=0 && !sqlite3Isspace(c) && c!=')' );
          if( c==')' ){
            i++;
          }else{
            *tokenType = TK_ILLEGAL;
          }
          break;
        }else if( c==':' && z[i+1]==':' ){
          i++;
#endif
        }else{
          break;
        }
      }
      if( n==0 ) *tokenType = TK_ILLEGAL;
      return i;
    }
    case 'x': case 'X': {
      if( z[1]=='\'' ){
        *tokenType = TK_BLOB;
        for(i=2; sqlite3Isxdigit(z[i]); i++){}
        if( z[i]!='\'' || i%2 ){
          *tokenType = TK_ILLEGAL;
          while( z[i] && z[i]!='\'' ){ i++; }
        }
        if( z[i] ) i++;
        return i;
      }
      /* Otherwise fall through to the next case */
    }
    default: {
      if( !IdChar(*z) ){
        break;
      }
      for(i=1; IdChar(z[i]); i++){}
      *tokenType = keywordCode((char*)z, i);
      return i;
    }
  }
  *tokenType = TK_ILLEGAL;
  return 1;
}

//	Run the parser on the given SQL string. The parser structure is passed in. An SQLITE_ status code is returned. If an error occurs then an and attempt is made to write an error message into memory obtained from sqlite3_malloc() and to make *pzErrMsg point to that error message.
func (pParse *Parse) Run(zSQL string, pzErrMsg *string) (nErr int) {
	db := pParse.db
	if db.activeVdbeCnt == 0 {
		db.u1.isInterrupted = false
	}
	pParse.rc = SQLITE_OK
	pParse.zTail = zSql
	assert( pzErrMsg != "" );
	pEngine := sqlite3ParserAlloc((void*(*)(size_t))sqlite3Malloc)
	if pEngine == nil {
		db.mallocFailed = true
		return SQLITE_NOMEM;
	}
	assert( pParse.pNewTable == nil && pParse.pNewTrigger == nil && pParse.nVar == 0 && pParse.nzVar == 0 && pParse.azVar == nil )
	enableLookaside := db.lookaside.bEnabled
	if db.lookaside.pStart != nil {
		db.lookaside.bEnabled = true
	}
	var tokenType		int				//	type of the next token
	lastTokenParsed := -1
	for i := 0; !db.mallocFailed && i < len(zSql); {
		assert( i >= 0 )
		pParse.sLastToken.z = zSql[i]
		pParse.sLastToken.n = sqlite3GetToken((unsigned char*)&zSql[i], &tokenType);
		i += pParse.sLastToken.n
		if i < 0 {
			pParse.rc = SQLITE_TOOBIG
			break
		}
		switch tokenType {
		case TK_SPACE:
			if db.u1.isInterrupted {
				pParse.SetErrorMsg("interrupt")
				pParse.rc = SQLITE_INTERRUPT
				goto abort_parse
			}
		case TK_ILLEGAL:
			*pzErrMsg = fmt.Sprintf("unrecognized token: \"%v\"", &pParse.sLastToken)
			nErr++
			goto abort_parse
		case TK_SEMI:
			pParse.zTail = &zSql[i]
			fallthrough
		default:
			sqlite3Parser(pEngine, tokenType, pParse.sLastToken, pParse);
			lastTokenParsed = tokenType
			if pParse.rc != SQLITE_OK {
				goto abort_parse
			}
		}
	}
abort_parse:
	if i == len(zSql) && nErr == 0 && pParse.rc == SQLITE_OK {
		if lastTokenParsed != TK_SEMI {
			sqlite3Parser(pEngine, TK_SEMI, pParse.sLastToken, pParse)
			pParse.zTail = &zSql[i]
		}
		sqlite3Parser(pEngine, 0, pParse.sLastToken, pParse)
	}
#ifdef YYTRACKMAXSTACKDEPTH
	sqlite3StatusSet(SQLITE_STATUS_PARSER_STACK, sqlite3ParserStackPeak(pEngine))
#endif /* YYDEBUG */
	sqlite3ParserFree(pEngine, sqlite3_free)
	db.lookaside.bEnabled = enableLookaside
	if db.mallocFailed {
		pParse.rc = SQLITE_NOMEM
	}
	if pParse.rc != SQLITE_OK && pParse.rc != SQLITE_DONE && pParse.zErrMsg == "" {
		pParse.zErrMsg = fmt.Sprintf("%v", sqlite3ErrStr(pParse.rc))
	}
	assert( pzErrMsg != "" )
	if pParse.zErrMsg != "" {
		*pzErrMsg = pParse.zErrMsg
		sqlite3_log(pParse.rc, "%s", *pzErrMsg)
		pParse.zErrMsg = ""
		nErr++
	}
	if pParse.pVdbe != nil && pParse.nErr > 0 && !pParse.nested {
		pParse.pVdbe.Delete()
		pParse.pVdbe = nil
	}
	if !pParse.nested {
		pParse.aTableLock = nil
		pParse.nTableLock = 0
	}
	pParse.apVtabLock = nil

	if !IN_DECLARE_VTAB {
		//	If the pParse.declareVtab flag is set, do not delete any table structure built up in pParse.pNewTable. The calling code (see vtab.c) will take responsibility for freeing the Table structure.
		db.DeleteTable(pParse.pNewTable)
	}

	db.DeleteTrigger(pParse.pNewTrigger)
	for i := pParse.nzVar - 1; i >= 0; i-- {
		pParse.azVar[i] = nil
	}
	pParse.azVar = nil
	pParse.aAlias = nil
	for pParse.pAinc != nil {
		p := pParse.pAinc
		pParse.pAinc = p.Next
	}
	for pParse.pZombieTab != nil {
		p := pParse.pZombieTab
		pParse.pZombieTab = p.NextZombie
		db.DeleteTable(p)
	}
	if nErr > 0 && pParse.rc == SQLITE_OK {
		pParse.rc = SQLITE_ERROR
	}
	return nErr
}