// Code generated from ./antlr4/V2Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // V2Parser

import "github.com/antlr4-go/antlr/v4"

// BaseV2ParserListener is a complete listener for a parse tree produced by V2Parser.
type BaseV2ParserListener struct{}

var _ V2ParserListener = &BaseV2ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseV2ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseV2ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseV2ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseV2ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseV2ParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseV2ParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterOpenBlock is called when production openBlock is entered.
func (s *BaseV2ParserListener) EnterOpenBlock(ctx *OpenBlockContext) {}

// ExitOpenBlock is called when production openBlock is exited.
func (s *BaseV2ParserListener) ExitOpenBlock(ctx *OpenBlockContext) {}

// EnterStructBlock is called when production structBlock is entered.
func (s *BaseV2ParserListener) EnterStructBlock(ctx *StructBlockContext) {}

// ExitStructBlock is called when production structBlock is exited.
func (s *BaseV2ParserListener) ExitStructBlock(ctx *StructBlockContext) {}

// EnterFuncBlock is called when production funcBlock is entered.
func (s *BaseV2ParserListener) EnterFuncBlock(ctx *FuncBlockContext) {}

// ExitFuncBlock is called when production funcBlock is exited.
func (s *BaseV2ParserListener) ExitFuncBlock(ctx *FuncBlockContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseV2ParserListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseV2ParserListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterDeclareBlock is called when production declareBlock is entered.
func (s *BaseV2ParserListener) EnterDeclareBlock(ctx *DeclareBlockContext) {}

// ExitDeclareBlock is called when production declareBlock is exited.
func (s *BaseV2ParserListener) ExitDeclareBlock(ctx *DeclareBlockContext) {}

// EnterFuncSig is called when production funcSig is entered.
func (s *BaseV2ParserListener) EnterFuncSig(ctx *FuncSigContext) {}

// ExitFuncSig is called when production funcSig is exited.
func (s *BaseV2ParserListener) ExitFuncSig(ctx *FuncSigContext) {}

// EnterFuncSignArgs is called when production funcSignArgs is entered.
func (s *BaseV2ParserListener) EnterFuncSignArgs(ctx *FuncSignArgsContext) {}

// ExitFuncSignArgs is called when production funcSignArgs is exited.
func (s *BaseV2ParserListener) ExitFuncSignArgs(ctx *FuncSignArgsContext) {}

// EnterType is called when production type is entered.
func (s *BaseV2ParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseV2ParserListener) ExitType(ctx *TypeContext) {}

// EnterVar is called when production var is entered.
func (s *BaseV2ParserListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BaseV2ParserListener) ExitVar(ctx *VarContext) {}

// EnterVars is called when production vars is entered.
func (s *BaseV2ParserListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BaseV2ParserListener) ExitVars(ctx *VarsContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseV2ParserListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseV2ParserListener) ExitIndex(ctx *IndexContext) {}

// EnterIndexes is called when production indexes is entered.
func (s *BaseV2ParserListener) EnterIndexes(ctx *IndexesContext) {}

// ExitIndexes is called when production indexes is exited.
func (s *BaseV2ParserListener) ExitIndexes(ctx *IndexesContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseV2ParserListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseV2ParserListener) ExitLambda(ctx *LambdaContext) {}

// EnterBinaryOper is called when production binaryOper is entered.
func (s *BaseV2ParserListener) EnterBinaryOper(ctx *BinaryOperContext) {}

// ExitBinaryOper is called when production binaryOper is exited.
func (s *BaseV2ParserListener) ExitBinaryOper(ctx *BinaryOperContext) {}

// EnterUnaryOper is called when production unaryOper is entered.
func (s *BaseV2ParserListener) EnterUnaryOper(ctx *UnaryOperContext) {}

// ExitUnaryOper is called when production unaryOper is exited.
func (s *BaseV2ParserListener) ExitUnaryOper(ctx *UnaryOperContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseV2ParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseV2ParserListener) ExitExpr(ctx *ExprContext) {}

// EnterExprWithLambda is called when production exprWithLambda is entered.
func (s *BaseV2ParserListener) EnterExprWithLambda(ctx *ExprWithLambdaContext) {}

// ExitExprWithLambda is called when production exprWithLambda is exited.
func (s *BaseV2ParserListener) ExitExprWithLambda(ctx *ExprWithLambdaContext) {}

// EnterExprOrAssign is called when production exprOrAssign is entered.
func (s *BaseV2ParserListener) EnterExprOrAssign(ctx *ExprOrAssignContext) {}

// ExitExprOrAssign is called when production exprOrAssign is exited.
func (s *BaseV2ParserListener) ExitExprOrAssign(ctx *ExprOrAssignContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseV2ParserListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseV2ParserListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncCallArgs is called when production funcCallArgs is entered.
func (s *BaseV2ParserListener) EnterFuncCallArgs(ctx *FuncCallArgsContext) {}

// ExitFuncCallArgs is called when production funcCallArgs is exited.
func (s *BaseV2ParserListener) ExitFuncCallArgs(ctx *FuncCallArgsContext) {}

// EnterStmtWithSep is called when production stmtWithSep is entered.
func (s *BaseV2ParserListener) EnterStmtWithSep(ctx *StmtWithSepContext) {}

// ExitStmtWithSep is called when production stmtWithSep is exited.
func (s *BaseV2ParserListener) ExitStmtWithSep(ctx *StmtWithSepContext) {}

// EnterOpenStmt is called when production openStmt is entered.
func (s *BaseV2ParserListener) EnterOpenStmt(ctx *OpenStmtContext) {}

// ExitOpenStmt is called when production openStmt is exited.
func (s *BaseV2ParserListener) ExitOpenStmt(ctx *OpenStmtContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseV2ParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseV2ParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralWithLambda is called when production literalWithLambda is entered.
func (s *BaseV2ParserListener) EnterLiteralWithLambda(ctx *LiteralWithLambdaContext) {}

// ExitLiteralWithLambda is called when production literalWithLambda is exited.
func (s *BaseV2ParserListener) ExitLiteralWithLambda(ctx *LiteralWithLambdaContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseV2ParserListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseV2ParserListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterMapInitializer is called when production mapInitializer is entered.
func (s *BaseV2ParserListener) EnterMapInitializer(ctx *MapInitializerContext) {}

// ExitMapInitializer is called when production mapInitializer is exited.
func (s *BaseV2ParserListener) ExitMapInitializer(ctx *MapInitializerContext) {}

// EnterMapPair is called when production mapPair is entered.
func (s *BaseV2ParserListener) EnterMapPair(ctx *MapPairContext) {}

// ExitMapPair is called when production mapPair is exited.
func (s *BaseV2ParserListener) ExitMapPair(ctx *MapPairContext) {}

// EnterStructInitializer is called when production structInitializer is entered.
func (s *BaseV2ParserListener) EnterStructInitializer(ctx *StructInitializerContext) {}

// ExitStructInitializer is called when production structInitializer is exited.
func (s *BaseV2ParserListener) ExitStructInitializer(ctx *StructInitializerContext) {}

// EnterStructElementInitializer is called when production structElementInitializer is entered.
func (s *BaseV2ParserListener) EnterStructElementInitializer(ctx *StructElementInitializerContext) {}

// ExitStructElementInitializer is called when production structElementInitializer is exited.
func (s *BaseV2ParserListener) ExitStructElementInitializer(ctx *StructElementInitializerContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseV2ParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseV2ParserListener) ExitStmt(ctx *StmtContext) {}

// EnterDeclareStmt is called when production declareStmt is entered.
func (s *BaseV2ParserListener) EnterDeclareStmt(ctx *DeclareStmtContext) {}

// ExitDeclareStmt is called when production declareStmt is exited.
func (s *BaseV2ParserListener) ExitDeclareStmt(ctx *DeclareStmtContext) {}

// EnterAssgnStmt is called when production assgnStmt is entered.
func (s *BaseV2ParserListener) EnterAssgnStmt(ctx *AssgnStmtContext) {}

// ExitAssgnStmt is called when production assgnStmt is exited.
func (s *BaseV2ParserListener) ExitAssgnStmt(ctx *AssgnStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseV2ParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseV2ParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterLoopStmt is called when production loopStmt is entered.
func (s *BaseV2ParserListener) EnterLoopStmt(ctx *LoopStmtContext) {}

// ExitLoopStmt is called when production loopStmt is exited.
func (s *BaseV2ParserListener) ExitLoopStmt(ctx *LoopStmtContext) {}

// EnterMatchStmt is called when production matchStmt is entered.
func (s *BaseV2ParserListener) EnterMatchStmt(ctx *MatchStmtContext) {}

// ExitMatchStmt is called when production matchStmt is exited.
func (s *BaseV2ParserListener) ExitMatchStmt(ctx *MatchStmtContext) {}

// EnterMatchCase is called when production matchCase is entered.
func (s *BaseV2ParserListener) EnterMatchCase(ctx *MatchCaseContext) {}

// ExitMatchCase is called when production matchCase is exited.
func (s *BaseV2ParserListener) ExitMatchCase(ctx *MatchCaseContext) {}

// EnterJumpStmt is called when production jumpStmt is entered.
func (s *BaseV2ParserListener) EnterJumpStmt(ctx *JumpStmtContext) {}

// ExitJumpStmt is called when production jumpStmt is exited.
func (s *BaseV2ParserListener) ExitJumpStmt(ctx *JumpStmtContext) {}

// EnterSep is called when production sep is entered.
func (s *BaseV2ParserListener) EnterSep(ctx *SepContext) {}

// ExitSep is called when production sep is exited.
func (s *BaseV2ParserListener) ExitSep(ctx *SepContext) {}
