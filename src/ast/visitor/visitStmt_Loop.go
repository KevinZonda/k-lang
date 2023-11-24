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
	panic("implement me")
}

func (v *AntlrVisitor) visitCStyleFor(ctx parser.ICStyleForContext) *node.CStyleFor {
	n := node.CStyleFor{
		Token:         token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:          v.visitCodeBlock(ctx.CodeBlock()),
		InitialExpr:   v.visitExpr(ctx.GetOnInit()),
		ConditionExpr: v.visitExpr(ctx.GetOnCondition()),
		AfterIterExpr: v.visitExpr(ctx.GetOnEnd()),
	}
	return &n
}

func (v *AntlrVisitor) visitWhileStyleFor(ctx parser.IWhileStyleForContext) *node.WhileStyleFor {
	n := node.WhileStyleFor{
		Token:         token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:          v.visitCodeBlock(ctx.CodeBlock()),
		ConditionExpr: v.visitExpr(ctx.Expr()),
	}
	return &n
}

func (v *AntlrVisitor) visitIterStyleFor(ctx parser.IIterForContext) *node.IterStyleFor {
	n := node.IterStyleFor{
		Token:    token.FromAntlrToken(ctx.For().GetSymbol()),
		Variable: v.visitIdentifier(ctx.Identifier()),
		Iterator: v.visitExpr(ctx.Expr()),
		Body:     v.visitCodeBlock(ctx.CodeBlock()),
		Type:     v.visitType(ctx.Type_()),
	}
	return &n
}
