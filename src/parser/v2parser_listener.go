// Code generated from ./antlr4/V2Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V2Parser

import "github.com/antlr4-go/antlr/v4"

// V2ParserListener is a complete listener for a parse tree produced by V2Parser.
type V2ParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStructBlock is called when entering the structBlock production.
	EnterStructBlock(c *StructBlockContext)

	// EnterFuncBlock is called when entering the funcBlock production.
	EnterFuncBlock(c *FuncBlockContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterDeclareBlock is called when entering the declareBlock production.
	EnterDeclareBlock(c *DeclareBlockContext)

	// EnterFuncSig is called when entering the funcSig production.
	EnterFuncSig(c *FuncSigContext)

	// EnterFuncReturnType is called when entering the funcReturnType production.
	EnterFuncReturnType(c *FuncReturnTypeContext)

	// EnterFuncSignArgs is called when entering the funcSignArgs production.
	EnterFuncSignArgs(c *FuncSignArgsContext)

	// EnterFuncSignArgItem is called when entering the funcSignArgItem production.
	EnterFuncSignArgItem(c *FuncSignArgItemContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterBaseVar is called when entering the baseVar production.
	EnterBaseVar(c *BaseVarContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterIndexes is called when entering the indexes production.
	EnterIndexes(c *IndexesContext)

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterUnaryOper is called when entering the unaryOper production.
	EnterUnaryOper(c *UnaryOperContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprWithLambda is called when entering the exprWithLambda production.
	EnterExprWithLambda(c *ExprWithLambdaContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncCallArgs is called when entering the funcCallArgs production.
	EnterFuncCallArgs(c *FuncCallArgsContext)

	// EnterOpenStmt is called when entering the openStmt production.
	EnterOpenStmt(c *OpenStmtContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterIdentifierPair is called when entering the identifierPair production.
	EnterIdentifierPair(c *IdentifierPairContext)

	// EnterMapPair is called when entering the mapPair production.
	EnterMapPair(c *MapPairContext)

	// EnterStructInitializer is called when entering the structInitializer production.
	EnterStructInitializer(c *StructInitializerContext)

	// EnterMapInitializer is called when entering the mapInitializer production.
	EnterMapInitializer(c *MapInitializerContext)

	// EnterCommaExpr is called when entering the commaExpr production.
	EnterCommaExpr(c *CommaExprContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterTypedIdentifiers is called when entering the typedIdentifiers production.
	EnterTypedIdentifiers(c *TypedIdentifiersContext)

	// EnterTypedIdentifier is called when entering the typedIdentifier production.
	EnterTypedIdentifier(c *TypedIdentifierContext)

	// EnterDeclareStmt is called when entering the declareStmt production.
	EnterDeclareStmt(c *DeclareStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterTryCatchSmt is called when entering the tryCatchSmt production.
	EnterTryCatchSmt(c *TryCatchSmtContext)

	// EnterCatchStmt is called when entering the catchStmt production.
	EnterCatchStmt(c *CatchStmtContext)

	// EnterLoopStmt is called when entering the loopStmt production.
	EnterLoopStmt(c *LoopStmtContext)

	// EnterCStyleFor is called when entering the cStyleFor production.
	EnterCStyleFor(c *CStyleForContext)

	// EnterIterFor is called when entering the iterFor production.
	EnterIterFor(c *IterForContext)

	// EnterWhileStyleFor is called when entering the whileStyleFor production.
	EnterWhileStyleFor(c *WhileStyleForContext)

	// EnterMatchStmt is called when entering the matchStmt production.
	EnterMatchStmt(c *MatchStmtContext)

	// EnterMatchCase is called when entering the matchCase production.
	EnterMatchCase(c *MatchCaseContext)

	// EnterJumpStmt is called when entering the jumpStmt production.
	EnterJumpStmt(c *JumpStmtContext)

	// EnterSep is called when entering the sep production.
	EnterSep(c *SepContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStructBlock is called when exiting the structBlock production.
	ExitStructBlock(c *StructBlockContext)

	// ExitFuncBlock is called when exiting the funcBlock production.
	ExitFuncBlock(c *FuncBlockContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitDeclareBlock is called when exiting the declareBlock production.
	ExitDeclareBlock(c *DeclareBlockContext)

	// ExitFuncSig is called when exiting the funcSig production.
	ExitFuncSig(c *FuncSigContext)

	// ExitFuncReturnType is called when exiting the funcReturnType production.
	ExitFuncReturnType(c *FuncReturnTypeContext)

	// ExitFuncSignArgs is called when exiting the funcSignArgs production.
	ExitFuncSignArgs(c *FuncSignArgsContext)

	// ExitFuncSignArgItem is called when exiting the funcSignArgItem production.
	ExitFuncSignArgItem(c *FuncSignArgItemContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitBaseVar is called when exiting the baseVar production.
	ExitBaseVar(c *BaseVarContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitIndexes is called when exiting the indexes production.
	ExitIndexes(c *IndexesContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitUnaryOper is called when exiting the unaryOper production.
	ExitUnaryOper(c *UnaryOperContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprWithLambda is called when exiting the exprWithLambda production.
	ExitExprWithLambda(c *ExprWithLambdaContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncCallArgs is called when exiting the funcCallArgs production.
	ExitFuncCallArgs(c *FuncCallArgsContext)

	// ExitOpenStmt is called when exiting the openStmt production.
	ExitOpenStmt(c *OpenStmtContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitIdentifierPair is called when exiting the identifierPair production.
	ExitIdentifierPair(c *IdentifierPairContext)

	// ExitMapPair is called when exiting the mapPair production.
	ExitMapPair(c *MapPairContext)

	// ExitStructInitializer is called when exiting the structInitializer production.
	ExitStructInitializer(c *StructInitializerContext)

	// ExitMapInitializer is called when exiting the mapInitializer production.
	ExitMapInitializer(c *MapInitializerContext)

	// ExitCommaExpr is called when exiting the commaExpr production.
	ExitCommaExpr(c *CommaExprContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitTypedIdentifiers is called when exiting the typedIdentifiers production.
	ExitTypedIdentifiers(c *TypedIdentifiersContext)

	// ExitTypedIdentifier is called when exiting the typedIdentifier production.
	ExitTypedIdentifier(c *TypedIdentifierContext)

	// ExitDeclareStmt is called when exiting the declareStmt production.
	ExitDeclareStmt(c *DeclareStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitTryCatchSmt is called when exiting the tryCatchSmt production.
	ExitTryCatchSmt(c *TryCatchSmtContext)

	// ExitCatchStmt is called when exiting the catchStmt production.
	ExitCatchStmt(c *CatchStmtContext)

	// ExitLoopStmt is called when exiting the loopStmt production.
	ExitLoopStmt(c *LoopStmtContext)

	// ExitCStyleFor is called when exiting the cStyleFor production.
	ExitCStyleFor(c *CStyleForContext)

	// ExitIterFor is called when exiting the iterFor production.
	ExitIterFor(c *IterForContext)

	// ExitWhileStyleFor is called when exiting the whileStyleFor production.
	ExitWhileStyleFor(c *WhileStyleForContext)

	// ExitMatchStmt is called when exiting the matchStmt production.
	ExitMatchStmt(c *MatchStmtContext)

	// ExitMatchCase is called when exiting the matchCase production.
	ExitMatchCase(c *MatchCaseContext)

	// ExitJumpStmt is called when exiting the jumpStmt production.
	ExitJumpStmt(c *JumpStmtContext)

	// ExitSep is called when exiting the sep production.
	ExitSep(c *SepContext)
}
