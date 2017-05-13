
grammar Nugget;

tokens {
	//IMAGINARY TOKENS
	ROOT,
	ATTR_LIST
}

nugget
   : load_stat EOF
   ;

load_stat
   : LOAD field FROM sourceidentifier ';'
   ;

field
   : Id (',' Id)*
   ;

sourceidentifier
   : Id
   ;

LOAD:'load';
SELECT: 'select';
FROM: 'from';

WS: [ \n\t\r]+ -> skip;

Id    : ('a'..'z' | 'A'..'Z' | '_') ('a'..'z' | 'A'..'Z' | '_' | Digit)*;
fragment Digit : '0'..'9';

