grammar Nugget;

tokens {
	ROOT,
	ATTR_LIST
}

nugget:   sin 		EOF
	| execute 	EOF
	| initextract 	EOF
	| printId	EOF
	| EOF
;

initextract : Id '=' target EXTRACT subtype
;

execute	: (Id '=') task sourceidentifier ';'
	| Id '|' task  ';'
;

subtype	: FILES
	| IMAGES
	| DOCUMENTS
	| ALL
; 

task    : HASH
	| EXTRACT
	| SELECT
;

target  : '"' Source '"' 
	| Source
;
	
field
	: Id (',' Id)*
	| '\'' field '\''
;

sourceidentifier
	: Id
	| '\'' sourceidentifier '\''
;

printId
	: Id
;

sin
	: SIN '(' NUMBER ')'
;

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

WS: [ \n\t\r]+ -> skip;

Id    	: ('a'..'z' | 'A'..'Z' | '_') ('a'..'z' | 'A'..'Z' | '_' | DIGIT)*;
Source  : ('a'..'z' | 'A'..'Z' | '_') ('a'..'z' | 'A'..'Z' | '_' | ':' | '.' | DIGIT)*;

NUMBER: DIGIT+ ;

fragment DIGIT : '0'..'9';
fragment DLMT  : ',';

