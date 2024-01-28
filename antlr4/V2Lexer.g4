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

Equals    : '==';
NotEq     : '!=';
GreaterEq : '>=';
LessEq    : '<=';
Greater   : '>';
Less      : '<';
Or        : '||' | 'or';
And       : '&&' | 'and';
Pow       : '^' | '**';
Mul       : '*';
Div       : '/';
Add       : '+';
Sub       : '-';
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

Try   : 'try';
Catch : 'catch';
If   : 'if';
Else : 'else';
For  : 'for';
Match : 'match';
Break : 'break';
Continue : 'next' | 'pass' | 'continue';

True : 'true' ;
False : 'false' ;

IntegerLiteral : [0-9]+ ;
NumberLiteral  : [0-9]+ ('.' [0-9]+)? ;
StringLiteral  : [$@]? '"' (~["] | '\\"')* '"' ;


Not       : '!' | 'not';
Assign    : ':=' | '=' | '<-';

Ref : '&';
Identifier: [a-zA-Z_][a-zA-Z_0-9]* ;
Comment   : '#' ~[\r\n]* -> channel(2); // channel 2 is for comments
WS: [ \t\n\r\f]+ -> skip ;

NewLine   : '\r'? '\n' | '\r';