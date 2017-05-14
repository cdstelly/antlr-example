grammar Nugget;

tokens {
	//IMAGINARY TOKENS
	ROOT,
	ATTR_LIST
}

nugget: sin EOF
	| execute EOF
	| initextract EOF
	| EOF
;

initextract : Id '=' target EXTRACT subtype
;

execute	: (Id '=') task sourceidentifier ';'
;

subtype	: FILES
	| IMAGES
;

task    : HASH
	| EXTRACT
;

target  : '"' Source '"' 
;
	
field
	: Id (',' Id)*
	| '\'' field '\''
;

sourceidentifier
	: Id
	| '\'' sourceidentifier '\''
;

sin
	: SIN '(' NUMBER ')'
;

SIN: 'sin';
LOAD:'load'|'LOAD';
FROM: 'from'|'FROM';

SELECT: 'select';

HASH: 'hash'; 
EXTRACT: 'extract';
FILES: 'files'|'FILES';
IMAGES: 'images'|'IMAGES';

WS: [ \n\t\r]+ -> skip;

Id    	: ('a'..'z' | 'A'..'Z' | '_') ('a'..'z' | 'A'..'Z' | '_' | DIGIT)*;
Source  : ('a'..'z' | 'A'..'Z' | '_') ('a'..'z' | 'A'..'Z' | '_' | ':' | '.' | DIGIT)*;

NUMBER: DIGIT+ ;

fragment DIGIT : '0'..'9';

