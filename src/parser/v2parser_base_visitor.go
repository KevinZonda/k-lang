// Code generated from .//antlr4//V2Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V2Parser

import "github.com/antlr4-go/antlr/v4"

type BaseV2ParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseV2ParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitStructBlock(ctx *StructBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitFuncBlock(ctx *FuncBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitCodeBlock(ctx *CodeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitDeclareBlock(ctx *DeclareBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitFuncSig(ctx *FuncSigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitFuncSignArgs(ctx *FuncSignArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitFuncSignArgItem(ctx *FuncSignArgItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitBaseVar(ctx *BaseVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitIndexes(ctx *IndexesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitUnaryOper(ctx *UnaryOperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitExprWithLambda(ctx *ExprWithLambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitFuncCallArgs(ctx *FuncCallArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitOpenStmt(ctx *OpenStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitIdentifierPair(ctx *IdentifierPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitMapPair(ctx *MapPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitStructInitializer(ctx *StructInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitMapInitializer(ctx *MapInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitCommaExpr(ctx *CommaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitVars(ctx *VarsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitTypedIdentifiers(ctx *TypedIdentifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitTypedIdentifier(ctx *TypedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitDeclareStmt(ctx *DeclareStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitTryCatchSmt(ctx *TryCatchSmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitCatchStmt(ctx *CatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitLoopStmt(ctx *LoopStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitCStyleFor(ctx *CStyleForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitIterFor(ctx *IterForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitWhileStyleFor(ctx *WhileStyleForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitMatchStmt(ctx *MatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitMatchCase(ctx *MatchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitJumpStmt(ctx *JumpStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseV2ParserVisitor) VisitSep(ctx *SepContext) interface{} {
	return v.VisitChildren(ctx)
}
