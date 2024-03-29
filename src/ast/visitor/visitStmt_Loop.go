package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitLoopStmt(ctx parser.ILoopStmtContext) node.Stmt {
	if ctx.CStyleFor() != nil {
		return v.visitCStyleFor(ctx.CStyleFor())
	}
	if ctx.WhileStyleFor() != nil {
		return v.visitWhileStyleFor(ctx.WhileStyleFor())
	}
	if ctx.IterFor() != nil {
		return v.visitIterStyleFor(ctx.IterFor())
	}

	v.appendErr(ctx, "unknown loop statement", nil)
	return nil
}

func (v *AntlrVisitor) visitCStyleFor(ctx parser.ICStyleForContext) *node.CStyleFor {
	sig := ctx.CStyleForSign()
	n := node.CStyleFor{
		Token:         token.FromAntlrToken(ctx.For().GetSymbol()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Body:          v.visitCodeBlock(ctx.CodeBlock()),
		InitialStmt:   v.visitStmt(sig.GetOnInit()),
		ConditionExpr: v.visitExpr(sig.GetOnCondition()),
		AfterIterStmt: v.visitStmt(sig.GetOnEnd()),
	}
	return &n
}

func (v *AntlrVisitor) visitWhileStyleFor(ctx parser.IWhileStyleForContext) *node.WhileStyleFor {
	n := node.WhileStyleFor{
		Token:         token.FromAntlrToken(ctx.For().GetSymbol()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Body:          v.visitCodeBlock(ctx.CodeBlock()),
		ConditionExpr: v.visitExpr(ctx.Expr()),
	}
	return &n
}

func (v *AntlrVisitor) visitIterStyleFor(ctx parser.IIterForContext) *node.IterStyleFor {
	sig := ctx.IterForSign()
	n := node.IterStyleFor{
		Token:    token.FromAntlrToken(ctx.For().GetSymbol()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Variable: v.visitIdentifier(sig.Identifier()),
		Iterator: v.visitExpr(sig.Expr()),
		Type:     v.visitType(sig.Type_()),
		Body:     v.visitCodeBlock(ctx.CodeBlock()),
	}
	return &n
}
