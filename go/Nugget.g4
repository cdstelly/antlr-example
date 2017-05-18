grammar Nugget; 

tokens {
	ROOT,
	ATTR_LIST
}

nugget:   sin 		EOF
	| execute 	EOF
	| initextract 	EOF
	| printId	EOF
	| assign	EOF
	| EOF
;

initextract : StrLit '=' target EXTRACT subtype
;

assign	: StrLit '=' StrLit ;

execute	: StrLit '=' StrLit '|' task 
	| StrLit '=' StrLit '|' task filter (',' filter)* 
	| StrLit '=' StrLit '|' task StrLit
;

filter  : filename
	| timefilter
;

filename : '"' StrLit '"'
;

timefilter: reference OPERATION date
;

reference : CTIME
	  | MTIME
;

date 	: '"' DATE '"' 
;

subtype	: FILES
	| IMAGES
	| DOCUMENTS
	| ALL
; 

task    : HASH
	| EXTRACT
	| SELECT
	| JOIN
;

target  : '"' StrLit '"' 
	| StrLit
;
	
field
	: StrLit (',' StrLit)*
	| '\'' field '\''
;

sourceidentifier
	: StrLit
	| '\'' sourceidentifier '\''
;

printId
	: StrLit
;

sin
	: SIN '(' NUMBER ')'
;

CTIME:  'ctime' | 'CTIME';
MTIME:  'mtime' | 'MTIME'; 
SIN: 	'sin';
LOAD:	'load'|'LOAD';
FROM: 	'from'|'FROM';
ALL: 	'all'|'ALL';
HASH: 	'hash'|'HASH';
SELECT: 'select'|'SELECT';
FILES: 	'files'|'FILES';
IMAGES:	'images'|'IMAGES';
DOCUMENTS: 	'documents' | 'DOCUMENTS'; 
EXTRACT:	'extract'|'EXTRACT';
JOIN: 'join'|'JOIN';

WS: [ \n\t\r]+ -> skip;


NUMBER: DIGIT+ ;
StrLit  : ('a'..'z' | 'A'..'Z' | '.' | ':' | '*' | DIGIT)+;
OPERATION : '>' | '<' | '<=' | '>=' | '=' ;
DATE	: DIGIT DIGIT '/' DIGIT DIGIT '/' DIGIT DIGIT DIGIT DIGIT;

fragment DIGIT : '0'..'9';
fragment DLMT  : ',';
fragment WILDCARD : '*';
