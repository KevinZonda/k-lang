// Code generated from .//antlr4//V2Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V2Parser

import "github.com/antlr4-go/antlr/v4"

// V2ParserListener is a complete listener for a parse tree produced by V2Parser.
type V2ParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterOpenBlock is called when entering the openBlock production.
	EnterOpenBlock(c *OpenBlockContext)

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

	// EnterBinaryOper is called when entering the binaryOper production.
	EnterBinaryOper(c *BinaryOperContext)

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

	// EnterStmtWithSep is called when entering the stmtWithSep production.
	EnterStmtWithSep(c *StmtWithSepContext)

	// EnterOpenStmt is called when entering the openStmt production.
	EnterOpenStmt(c *OpenStmtContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralWithLambda is called when entering the literalWithLambda production.
	EnterLiteralWithLambda(c *LiteralWithLambdaContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterMapInitializer is called when entering the mapInitializer production.
	EnterMapInitializer(c *MapInitializerContext)

	// EnterMapPair is called when entering the mapPair production.
	EnterMapPair(c *MapPairContext)

	// EnterStructInitializer is called when entering the structInitializer production.
	EnterStructInitializer(c *StructInitializerContext)

	// EnterStructElementInitializer is called when entering the structElementInitializer production.
	EnterStructElementInitializer(c *StructElementInitializerContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDeclareStmt is called when entering the declareStmt production.
	EnterDeclareStmt(c *DeclareStmtContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

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

	// ExitOpenBlock is called when exiting the openBlock production.
	ExitOpenBlock(c *OpenBlockContext)

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

	// ExitBinaryOper is called when exiting the binaryOper production.
	ExitBinaryOper(c *BinaryOperContext)

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

	// ExitStmtWithSep is called when exiting the stmtWithSep production.
	ExitStmtWithSep(c *StmtWithSepContext)

	// ExitOpenStmt is called when exiting the openStmt production.
	ExitOpenStmt(c *OpenStmtContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralWithLambda is called when exiting the literalWithLambda production.
	ExitLiteralWithLambda(c *LiteralWithLambdaContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitMapInitializer is called when exiting the mapInitializer production.
	ExitMapInitializer(c *MapInitializerContext)

	// ExitMapPair is called when exiting the mapPair production.
	ExitMapPair(c *MapPairContext)

	// ExitStructInitializer is called when exiting the structInitializer production.
	ExitStructInitializer(c *StructInitializerContext)

	// ExitStructElementInitializer is called when exiting the structElementInitializer production.
	ExitStructElementInitializer(c *StructElementInitializerContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDeclareStmt is called when exiting the declareStmt production.
	ExitDeclareStmt(c *DeclareStmtContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

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
