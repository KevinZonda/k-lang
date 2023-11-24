package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitLoopStmt(ctx *parser.LoopStmtContext) node.Stmt {
	if ctx.CStyleFor() != nil {
		return v.visitCStyleFor(ctx.CStyleFor().(*parser.CStyleForContext))
	}
	if ctx.WhileStyleFor() != nil {
		return v.visitWhileStyleFor(ctx.WhileStyleFor().(*parser.WhileStyleForContext))
	}
	if ctx.IterFor() != nil {
		return v.visitIterStyleFor(ctx.IterFor().(*parser.IterForContext))
	}
	panic("implement me")
}

func (v *AntlrVisitor) visitCStyleFor(ctx *parser.CStyleForContext) *node.CStyleFor {
	n := node.CStyleFor{
		Token:         token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:          v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)),
		InitialExpr:   v.visitExpr(toPtr[parser.ExprContext](ctx.GetOnInit())),
		ConditionExpr: v.visitExpr(toPtr[parser.ExprContext](ctx.GetOnCondition())),
		AfterIterExpr: v.visitExpr(toPtr[parser.ExprContext](ctx.GetOnEnd())),
	}
	return &n
}

func (v *AntlrVisitor) visitWhileStyleFor(ctx *parser.WhileStyleForContext) *node.WhileStyleFor {
	n := node.WhileStyleFor{
		Token: token.FromAntlrToken(ctx.For().GetSymbol()),
		Body:  v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)),
	}
	if ctx.Expr() != nil {
		n.ConditionExpr = v.visitExpr(ctx.Expr().(*parser.ExprContext))
	}
	return &n
}

func (v *AntlrVisitor) visitIterStyleFor(ctx *parser.IterForContext) *node.IterStyleFor {
	n := node.IterStyleFor{
		Token:    token.FromAntlrToken(ctx.For().GetSymbol()),
		Variable: v.visitIdentifier(ctx.Identifier()),
		Iterator: v.visitExpr(ctx.Expr().(*parser.ExprContext)),
		Body:     v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)),
	}
	if ctx.Type_() != nil {
		t := v.VisitType(toPtr[parser.TypeContext](ctx.Type_()))
		if t != nil {
			n.Type = t.(*node.Type)
		}
	}
	return &n
}
