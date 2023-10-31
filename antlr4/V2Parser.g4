parser grammar V2Parser;
options { tokenVocab=V2Lexer; }

program
    : openBlock ((structBlock | funcBlock | stmt | expr) (sep*))* EOF
    ;

openBlock : openStmt*;
structBlock
    : Struct Identifier declareBlock
    ;

funcBlock
    : Function funcSig codeBlock
    ;

codeBlock : LBrack (stmt sep*)* RBrack;
declareBlock : LBrack declareStmt* RBrack;
funcSig : Identifier LParen funcSignArgs? RParen type?;
funcSignArgs : type Identifier (Comma type Identifier)*;

type: Map | Function | Identifier (LSquare RSquare)?;
var : Identifier (Dot Identifier)*;
vars
    : var (Comma vars)*
    ;

index : LSquare expr RSquare;
indexes : index+;

lambda : LParen funcSignArgs RParen type? codeBlock;

binaryOper : Equals | NotEq | Greater | Less | GreaterEq | LessEq | (Or | And) | Pow | (Mul | Div) | Mod | (Add | Sub);
unaryOper : Add | Sub | Not;
expr
    : funcCall
    | unaryOper expr
    | expr binaryOper expr
    | literal
    | LParen expr RParen
    | Identifier
    | expr indexes
    ;

exprWithLambda : lambda | expr;
exprOrAssign : expr | assgnStmt;

funcCall : var LParen funcCallArgs? RParen ;
funcCallArgs : expr (Comma expr)*;

stmtWithSep : stmt sep*;

openStmt : Open StringLiteral (As Identifier)?;

literal : True | False | IntegerLiteral | NumberLiteral | StringLiteral | arrayInitializer | mapInitializer | structInitializer;
literalWithLambda : literal | lambda;

arrayInitializer : type? LSquare (expr Comma?)* RSquare;
mapInitializer : Map LBrack mapPair* RBrack;
mapPair : LParen expr Comma exprWithLambda RParen | expr To exprWithLambda;
structInitializer : type LBrack (structElementInitializer Comma?)* RBrack;
structElementInitializer : Identifier Comma expr;

stmt
    : declareStmt
    | assgnStmt
    | jumpStmt
    | ifStmt
    | loopStmt
    | matchStmt
    | codeBlock
    | funcCall
    ;

declareStmt
    : type vars;

assgnStmt
    : type? var indexes? Assign exprWithLambda
    ;

ifStmt
    : If LParen expr RParen codeBlock (Else codeBlock)?
    ;

loopStmt
    : For LParen exprOrAssign? Semi expr? Semi expr? RParen codeBlock
    | For LParen expr? RParen codeBlock
    ;

matchStmt
    : Match LParen expr RParen LBrack matchCase* RBrack
    ;

matchCase
    : Case expr Col codeBlock
    | Default Col codeBlock
    ;


jumpStmt
    : Continue
    | Break
    | Return expr?
    ;

sep : NewLine | Semi;