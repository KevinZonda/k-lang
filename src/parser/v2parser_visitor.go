// Code generated from ./antlr4/V2Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V2Parser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by V2Parser.
type V2ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by V2Parser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by V2Parser#openBlock.
	VisitOpenBlock(ctx *OpenBlockContext) interface{}

	// Visit a parse tree produced by V2Parser#structBlock.
	VisitStructBlock(ctx *StructBlockContext) interface{}

	// Visit a parse tree produced by V2Parser#funcBlock.
	VisitFuncBlock(ctx *FuncBlockContext) interface{}

	// Visit a parse tree produced by V2Parser#codeBlock.
	VisitCodeBlock(ctx *CodeBlockContext) interface{}

	// Visit a parse tree produced by V2Parser#declareBlock.
	VisitDeclareBlock(ctx *DeclareBlockContext) interface{}

	// Visit a parse tree produced by V2Parser#funcSig.
	VisitFuncSig(ctx *FuncSigContext) interface{}

	// Visit a parse tree produced by V2Parser#funcSignArgs.
	VisitFuncSignArgs(ctx *FuncSignArgsContext) interface{}

	// Visit a parse tree produced by V2Parser#funcSignArgItem.
	VisitFuncSignArgItem(ctx *FuncSignArgItemContext) interface{}

	// Visit a parse tree produced by V2Parser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by V2Parser#var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by V2Parser#baseVar.
	VisitBaseVar(ctx *BaseVarContext) interface{}

	// Visit a parse tree produced by V2Parser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by V2Parser#indexes.
	VisitIndexes(ctx *IndexesContext) interface{}

	// Visit a parse tree produced by V2Parser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by V2Parser#unaryOper.
	VisitUnaryOper(ctx *UnaryOperContext) interface{}

	// Visit a parse tree produced by V2Parser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by V2Parser#exprWithLambda.
	VisitExprWithLambda(ctx *ExprWithLambdaContext) interface{}

	// Visit a parse tree produced by V2Parser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by V2Parser#funcCallArgs.
	VisitFuncCallArgs(ctx *FuncCallArgsContext) interface{}

	// Visit a parse tree produced by V2Parser#stmtWithSep.
	VisitStmtWithSep(ctx *StmtWithSepContext) interface{}

	// Visit a parse tree produced by V2Parser#openStmt.
	VisitOpenStmt(ctx *OpenStmtContext) interface{}

	// Visit a parse tree produced by V2Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by V2Parser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by V2Parser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by V2Parser#identifierPair.
	VisitIdentifierPair(ctx *IdentifierPairContext) interface{}

	// Visit a parse tree produced by V2Parser#mapPair.
	VisitMapPair(ctx *MapPairContext) interface{}

	// Visit a parse tree produced by V2Parser#structInitializer.
	VisitStructInitializer(ctx *StructInitializerContext) interface{}

	// Visit a parse tree produced by V2Parser#mapInitializer.
	VisitMapInitializer(ctx *MapInitializerContext) interface{}

	// Visit a parse tree produced by V2Parser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by V2Parser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by V2Parser#declareStmt.
	VisitDeclareStmt(ctx *DeclareStmtContext) interface{}

	// Visit a parse tree produced by V2Parser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by V2Parser#loopStmt.
	VisitLoopStmt(ctx *LoopStmtContext) interface{}

	// Visit a parse tree produced by V2Parser#cStyleFor.
	VisitCStyleFor(ctx *CStyleForContext) interface{}

	// Visit a parse tree produced by V2Parser#iterFor.
	VisitIterFor(ctx *IterForContext) interface{}

	// Visit a parse tree produced by V2Parser#whileStyleFor.
	VisitWhileStyleFor(ctx *WhileStyleForContext) interface{}

	// Visit a parse tree produced by V2Parser#matchStmt.
	VisitMatchStmt(ctx *MatchStmtContext) interface{}

	// Visit a parse tree produced by V2Parser#matchCase.
	VisitMatchCase(ctx *MatchCaseContext) interface{}

	// Visit a parse tree produced by V2Parser#jumpStmt.
	VisitJumpStmt(ctx *JumpStmtContext) interface{}

	// Visit a parse tree produced by V2Parser#sep.
	VisitSep(ctx *SepContext) interface{}
}
