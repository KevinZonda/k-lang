package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitMathStmt(ctx *parser.MatchStmtContext) any {
	match := &node.MatchStmt{
		Token: token.FromAntlrToken(ctx.Match().GetSymbol()),
		Match: v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr),
	}
	for _, c := range ctx.AllMatchCase() {
		m := c.(*parser.MatchCaseContext)
		if m.Default() == nil {
			match.Cases = append(match.Cases, v.VisitMatchCase(m).(*node.MatchCase))
		} else {
			match.Default = typeCastToPtr[node.CodeBlock](v.VisitCodeBlock(typeCastToPtr[parser.CodeBlockContext](m.CodeBlock())))
		}
	}
	return match
}

func (v *AntlrVisitor) VisitMatchCase(ctx *parser.MatchCaseContext) any {
	return &node.MatchCase{
		Expr: v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr),
		Body: v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
}
