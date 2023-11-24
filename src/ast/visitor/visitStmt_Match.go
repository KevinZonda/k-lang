package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitMathStmt(ctx *parser.MatchStmtContext) *node.MatchStmt {
	match := &node.MatchStmt{
		Token: token.FromAntlrToken(ctx.Match().GetSymbol()),
		Match: v.visitExpr(ctx.Expr().(*parser.ExprContext)),
	}
	for _, c := range ctx.AllMatchCase() {
		m := c.(*parser.MatchCaseContext)
		if m.Default() == nil {
			match.Cases = append(match.Cases, v.visitMatchCase(m))
		} else {
			match.Default = v.visitCodeBlock(toPtr[parser.CodeBlockContext](m.CodeBlock()))
		}
	}
	return match
}

func (v *AntlrVisitor) visitMatchCase(ctx *parser.MatchCaseContext) *node.MatchCase {
	return &node.MatchCase{
		Expr: v.visitExpr(ctx.Expr().(*parser.ExprContext)),
		Body: v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)),
	}
}
