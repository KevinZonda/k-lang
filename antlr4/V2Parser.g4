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
funcSignArgs : funcSignArgItem? (Comma funcSignArgItem)*;
funcSignArgItem : type? Identifier;

type: Map | Function | Identifier (LSquare RSquare)?;
var : baseVar (Dot baseVar)*; // x.y[X].z
baseVar : Identifier indexes?;

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
exprOrAssign : expr | assignStmt;

funcCall : var LParen funcCallArgs? RParen ;
funcCallArgs : expr (Comma expr)*;

stmtWithSep : stmt sep*;

openStmt : Open StringLiteral (As Identifier)?;

literal : True | False | IntegerLiteral | NumberLiteral | StringLiteral | arrayInitializer | mapInitializer | structInitializer;
literalWithLambda : literal | lambda;

arrayInitializer : type? LSquare (expr Comma?)* RSquare;
mapInitializer : Map LBrack (mapPair (Comma*))* RBrack;
mapPair : expr (Col | To) exprWithLambda | LParen mapPair RParen;
structInitializer : type LBrack (structElementInitializer Comma?)* RBrack;
structElementInitializer : Identifier Comma expr;

stmt
    : declareStmt
    | assignStmt
    | jumpStmt
    | ifStmt
    | loopStmt
    | matchStmt
    | codeBlock
    | funcCall
    ;

declareStmt
    : type Identifier (Comma Identifier)*;

assignStmt
    : type? var Assign exprWithLambda
    ;

ifStmt
    : If LParen expr RParen codeBlock (Else codeBlock)?
    ;

loopStmt
    : iterFor
    | cStyleFor
    | whileStyleFor
    ;

cStyleFor : For LParen onInit=exprOrAssign? Semi onCondition=expr? Semi onEnd=expr? RParen codeBlock;
iterFor : For LParen type? Identifier Col expr RParen codeBlock;
whileStyleFor : For (LParen expr? RParen)? codeBlock;

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