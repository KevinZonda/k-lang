parser grammar V2Parser;
options { tokenVocab=V2Lexer; }

program
    : openBlock (structBlock | funcBlock | stmt)* EOF
    ;

openBlock : openStmt*;
structBlock
    : Struct Identity declareBlock
    ;

funcBlock
    : Function funcSig codeBlock
    ;

codeBlock : LBrack stmtWithSep* RBrack;
declareBlock : LBrack declareStmt* RBrack;
funcSig : Identity LParen funcSignArgs? RParen type?;
funcSignArgs : type Identity (Comma type Identity)*;

type: Map | Function | Identity (LSquare RSquare)?;
var : Identity (Dot Identity)*;
vars
    : var (Comma vars)*
    ;
varWithIdx : var indexes;

index : LSquare expr RSquare;
indexes : index*;

lambda : LParen funcSignArgs RParen type? codeBlock;

expr
    : funcCall
    | Not expr
    | expr BinaryOper expr
    | literal
    | LParen expr RParen
    | Identity
    ;

exprWithLambda : lambda | expr;
exprOrAssign : expr | assgnStmt;

funcCall : var LParen funcCallArgs? RParen ;
funcCallArgs : expr (Comma expr)*;

stmtWithSep : stmt sep*;

openStmt : Open StringLiteral (As Identity)?;

literal : IntegerLiteral | NumberLiteral | StringLiteral | arrayInitializer | mapInitializer | structInitializer;
literalWithLambda : literal | lambda;

arrayInitializer : type? LSquare (expr Comma?)* RSquare;
mapInitializer : Map LBrack mapPair* RBrack;
mapPair : LParen expr Comma exprWithLambda RParen | expr To exprWithLambda;
structInitializer : type LBrack (structElementInitializer Comma?)* RBrack;
structElementInitializer : Identity Comma expr;

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
    : type? varWithIdx Assign exprWithLambda
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