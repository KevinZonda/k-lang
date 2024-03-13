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
			match.Default = v.visitCaseBody(m.CaseBlock())
		}
	}
	return match
}

func (v *AntlrVisitor) visitMatchCase(ctx parser.IMatchCaseContext) *node.MatchCase {
	return &node.MatchCase{
		Expr: v.visitExpr(ctx.GetCondition()),
		Body: v.visitCaseBody(ctx.CaseBlock()),
	}
}

func (v *AntlrVisitor) visitCaseBody(caseBlock parser.ICaseBlockContext) *node.CodeBlock {
	if caseBlock == nil {
		return nil
	}
	if caseBlock.CodeBlock() != nil {
		return v.visitCodeBlock(caseBlock.CodeBlock())
	}
	var body []node.Node
	for _, m := range caseBlock.GetChildren() {
		switch m.(type) {
		case *parser.StmtContext:
			body = append(body, v.visitStmt(m.(*parser.StmtContext)))
		case *parser.ExprContext:
			body = append(body, v.visitExpr(m.(*parser.ExprContext)))
		}
	}
	return &node.CodeBlock{
		Nodes: body,
		Token: token.FromAntlrToken(caseBlock.GetStart()).WithEnd(caseBlock.GetStop()),
	}
	return nil

}
