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

codeBlock : LBrack ((stmt | expr) sep*)* RBrack;
declareBlock : LBrack declareStmt* RBrack;
funcSig : Identifier LParen funcSignArgs RParen type?;
funcSignArgs : funcSignArgItem? (Comma funcSignArgItem)*;
funcSignArgItem : type? Identifier;

type: Map | Function | (PackageName=Identifier Dot)? TypeName=Identifier (LSquare RSquare)?;
var : baseVar (Dot baseVar)*; // x.y[X].z
baseVar : Identifier indexes?;

index : LSquare expr RSquare;
indexes : index+;

lambda : Function LParen funcSignArgs RParen type? codeBlock;

//boolExpr
//    : expr (Equals | NotEq | Greater | Less | GreaterEq | LessEq) expr
//    | expr (Or | And) expr
//    ;
//arithExpr
//    : expr Pow expr
//    | expr (Mul | Div | Mod) expr
//    | expr (Add | Sub) expr
//    ;
unaryOper : Add | Sub | Not;
expr
    : funcCall
    | unaryOper expr
    | literal
    | LParen expr RParen
    | LHS=expr OP=(Equals | NotEq | Greater | Less | GreaterEq | LessEq) RHS=expr
    | LHS=expr OP=(Or | And)        RHS=expr
    | LHS=expr OP=Pow               RHS=expr
    | LHS=expr OP=(Mul | Div | Mod) RHS=expr
    | LHS=expr OP=(Add | Sub)       RHS=expr
    | Identifier
    | expr indexes
    | LHS=expr Dot RHS=expr
    | assignStmt
    ;

exprWithLambda : lambda | expr;

funcCall : Identifier LParen funcCallArgs? RParen ;
funcCallArgs : expr (Comma expr)*;

stmtWithSep : stmt sep*;

openStmt : Open StringLiteral (As Identifier)?;

literal : True | False | IntegerLiteral | NumberLiteral | StringLiteral | arrayInitializer | mapInitializer | structInitializer;
literalWithLambda : literal | lambda;

arrayInitializer : type? LSquare (expr Comma?)* RSquare;
mapInitializer : Map? LBrack (mapPair (Comma*))* RBrack;
mapPair : expr (Col | To) exprWithLambda | LParen mapPair RParen;
structInitializer : type LBrack (structElementInitializer Comma?)* RBrack;
structElementInitializer : Identifier Comma expr;

stmt
    : assignStmt
    | declareStmt
    | jumpStmt
    | ifStmt
    | loopStmt
    | matchStmt
    | codeBlock
    | funcCall
    ;
assignStmt
    : type? var Assign exprWithLambda ;

declareStmt
    : type Identifier (Comma Identifier)*;


ifStmt
    : If LParen expr RParen codeBlock (Else codeBlock)?
    ;

loopStmt
    : iterFor
    | cStyleFor
    | whileStyleFor
    ;

cStyleFor : For LParen onInit=expr? Semi onCondition=expr? Semi onEnd=expr? RParen codeBlock;
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
    | Return exprWithLambda?
    ;

sep : NewLine | Semi;