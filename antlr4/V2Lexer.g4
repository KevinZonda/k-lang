// DELETE THIS CONTENT IF YOU PUT COMBINED GRAMMAR IN Parser TAB
lexer grammar V2Lexer;

LBrack  : '{';
RBrack  : '}';
LParen  : '(';
RParen  : ')';
LSquare : '[';
RSquare : ']';
Comma   : ',';
Col     : ':';
Semi    : ';';

BinaryOper : Equals | NotEq | Greater | Less | GreaterEq | LessEq | Or | And | Add | Sub | Mul | Div | Mod;
Equals    : '==';
NotEq     : '!=';
GreaterEq : '>=';
LessEq    : '<=';
Greater   : '>';
Less      : '<';
Or        : '||' | 'or';
And       : '&&' | 'and';
Add       : '+';
Sub       : '-';
Mul       : '*';
Div       : '/';
Mod       : '%' | 'mod';
Dot       : '.';

To : '->';

Struct   : 'struct';
Map      : 'map';
Function : 'fn';

Return  : 'return';
Case    : 'case';
Default : 'default';
Open    : 'open';
As      : 'as';

If   : 'if';
Else : 'else';
For  : 'for';
Match : 'match';
Break : 'break';
Continue : 'continue';

IntegerLiteral : ('-')? [0-9]+ ;
NumberLiteral  : ('-')? [0-9]+ ('.' [0-9]+)? ;
StringLiteral  : '"' .*? '"' ;


Not       : '!' | 'not';
Assign    : ':=';
Identity: [a-zA-Z_][a-zA-Z_0-9]* ;
WS: [ \t\n\r\f]+ -> skip ;

NewLine   : '\r'? '\n' | '\r';