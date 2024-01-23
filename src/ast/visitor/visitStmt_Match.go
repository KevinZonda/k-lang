package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitMathStmt(ctx parser.IMatchStmtContext) *node.MatchStmt {
	match := &node.MatchStmt{
		Token: token.FromAntlrToken(ctx.Match().GetSymbol()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Match: v.visitExpr(ctx.Expr()),
	}
	for _, c := range ctx.AllMatchCase() {
		m := c.(*parser.MatchCaseContext)
		if m.Default() == nil {
			match.Cases = append(match.Cases, v.visitMatchCase(m))
		} else {
			match.Default = v.visitCodeBlock(m.CodeBlock())
		}
	}
	return match
}

func (v *AntlrVisitor) visitMatchCase(ctx parser.IMatchCaseContext) *node.MatchCase {
	return &node.MatchCase{
		Expr: v.visitExpr(ctx.Expr()),
		Body: v.visitCodeBlock(ctx.CodeBlock()),
	}
}
