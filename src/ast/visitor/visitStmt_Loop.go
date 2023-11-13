package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitLoopStmt(ctx *parser.LoopStmtContext) interface{} {
	if ctx.CStyleFor() != nil {
		return v.VisitCStyleFor(ctx.CStyleFor().(*parser.CStyleForContext))
	}
	if ctx.WhileStyleFor() != nil {
		return v.VisitWhileStyleFor(ctx.WhileStyleFor().(*parser.WhileStyleForContext))
	}
	if ctx.IterFor() != nil {
		return v.VisitIterStyleFor(ctx.IterFor().(*parser.IterForContext))
	}
	panic("implement me")
}

func (v *AntlrVisitor) VisitCStyleFor(ctx *parser.CStyleForContext) any {
	n := node.CStyleFor{
		Token: token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:  v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
	if ctx.GetOnInit() != nil {
		n.InitialExpr = v.VisitExpr(ctx.GetOnInit().(*parser.ExprContext)).(node.Node)
	}

	if ctx.GetOnCondition() != nil {
		n.ConditionExpr = v.VisitExpr(ctx.GetOnCondition().(*parser.ExprContext)).(node.Expr)
	}

	if ctx.GetOnEnd() != nil {
		n.AfterIterExpr = v.VisitExpr(ctx.GetOnEnd().(*parser.ExprContext)).(node.Expr)
	}

	return &n
}

func (v *AntlrVisitor) VisitWhileStyleFor(ctx *parser.WhileStyleForContext) any {
	n := node.WhileStyleFor{
		Token: token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:  v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
	if ctx.Expr() != nil {
		n.ConditionExpr = v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr)
	}
	return &n
}

func (v *AntlrVisitor) VisitIterStyleFor(ctx *parser.IterForContext) any {
	n := node.IterStyleFor{
		Token:    token.FromAntlrToken(ctx.For().GetSymbol()),
		Variable: v.visitIdentifier(ctx.Identifier()),
		Iterator: v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr),
		Body:     v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
	if ctx.Type_() != nil {
		t := v.VisitType(typeCastToPtr[parser.TypeContext](ctx.Type_()))
		if t != nil {
			n.Type = t.(*node.Identifier)
		}
	}
	return &n
}
