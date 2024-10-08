parser grammar V2Parser;
options { tokenVocab=V2Lexer; }

program
    : ((structBlock | funcBlock | stmt | commaExpr) (sep*))* EOF
    ;

structBlock
    : Struct Identifier declareBlock
    ;

funcBlock
    : Function funcSig codeBlock
    ;

codeBlock : LBrack ((stmt | expr) sep*)* RBrack;
declareBlock : LBrack declareStmt* RBrack;
funcSig : Identifier LParen funcSignArgs RParen funcReturnType?;
funcReturnType :  (Flag=(To|Col)? (type|LParen type (Comma type)* RParen));
funcSignArgs : funcSignArgItem? (Comma funcSignArgItem)*;
funcSignArgItem : type? Ref? Identifier | Ref? Identifier (Col type)?;

type: Map | Function | (PackageName=Identifier Dot)? TypeName=Identifier (LSquare RSquare)? Question?;
var : baseVar (Dot baseVar)*; // x.y[X].z
baseVar : Identifier indexes?;

index : LSquare expr RSquare;
indexes : index+;

lambda : Function LParen funcSignArgs RParen funcReturnType? codeBlock
       | LParen funcSignArgs RParen funcReturnType? To (codeBlock | exprWithLambda);

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
    : Ref expr
    | LHS=expr Dot RHS=expr
    | LHS=expr LSquare Index=expr Col? EndIndex=expr? RSquare
    | funcCall
    | CallExpr=expr LParen funcCallArgs? RParen
    | unaryOper expr
    | LParen expr RParen
    | LHS=expr OP=Pow               RHS=expr
    | LHS=expr OP=(Mul | Div | Mod) RHS=expr
    | LHS=expr OP=(Add | Sub)       RHS=expr
    | LHS=expr OP=(Equals | NotEq | Greater | Less | GreaterEq | LessEq) RHS=expr
    | LHS=expr OP=(Or | And)        RHS=expr
    | structInitializer
    | Identifier
    | literal
    | initializer
    | expr indexes
    | TriCond=expr Question IfTrue=expr Col IfFalse=expr
    ;

exprWithLambda : lambda | expr
               | TriCond=expr Question IfTrue=exprWithLambda Col IfFalse=exprWithLambda
               | CallExpr=exprWithLambda LParen funcCallArgs? RParen
               | LParen ParenExpr=exprWithLambda RParen;

funcCall : Identifier LParen funcCallArgs? RParen ;
funcCallArgs : exprWithLambda (Comma exprWithLambda)*;

openStmt : Open StringLiteral (As AS=(Identifier | Dot))?;

literal : Nil | True | False | IntegerLiteral | NumberLiteral | StringLiteral;
initializer : arrayInitializer | mapInitializer;

arrayInitializer : type? LSquare (expr Comma?)* RSquare;
identifierPair : LHS=Identifier  (Col | To) RHS=exprWithLambda;
mapPair : LHS=expr (Col | To) RHS=exprWithLambda | LParen mapPair RParen | identifierPair;
structInitializer : (type | Struct) LBrack (identifierPair Comma?)* RBrack;
// structElementInitializer : Identifier Comma expr;
mapInitializer : Map? LBrack (mapPair (Comma*))* RBrack;
commaExpr : exprWithLambda (Comma exprWithLambda)*;

stmt
    : assignStmt
    | declareStmt
    | jumpStmt
    | ifStmt
    | loopStmt
    | matchStmt
    | uOperStmt
    | codeBlock
    | tryCatchSmt
    | openStmt
    | funcCall
    ;

uOperStmt : Identifier (AddAdd | SubSub);

assignStmt
    : type? var Assign exprWithLambda
    | var Col type Assign exprWithLambda
    | vars Assign commaExpr
    ;

vars : var (Comma var)*;

typedIdentifiers : type Identifier (Comma Identifier)* | Identifier (Comma Identifier)* Col type;
typedIdentifier : type? Identifier | Identifier (Col type)?;
declareStmt
    : typedIdentifiers
    | typedIdentifier Assign exprWithLambda
    | funcBlock
    ;


ifStmt
    : If expr codeBlock (Else (codeBlock|ifStmt))?
    ;

tryCatchSmt
    : Try codeBlock catchStmt
    ;
catchStmt
    : Catch (LParen type? Identifier RParen)? codeBlock
    ;

loopStmt
    : iterFor
    | cStyleFor
    | whileStyleFor
    ;

cStyleFor : For ((LParen cStyleForSign RParen) | cStyleForSign) codeBlock;
cStyleForSign : onInit=stmt? Semi onCondition=expr? Semi onEnd=stmt?;
iterFor  : For  ((LParen iterForSign RParen)   | iterForSign) codeBlock;
iterForSign : type? Identifier Col expr;
whileStyleFor : For expr? codeBlock;

matchStmt
    : Match expr LBrack matchCase* RBrack
    ;

matchCase
    : Case expr (Comma expr)* caseBlock
    | Default caseBlock
    ;

caseBlock : codeBlock |  Col ((stmt | expr) sep*)*;


jumpStmt
    : Continue
    | Break
    | Return (exprWithLambda (Comma exprWithLambda)*)?
    ;

sep : NewLine | Semi;