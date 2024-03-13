package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

func (e *Eval) EvalMatchStmt(m *node.MatchStmt) {
	e.currentToken = m.GetToken()
	match := e.EvalExpr(m.Match).EnsureValue()
	for _, c := range m.Cases {
		for _, expr := range c.Expr {
			if e.EvalExpr(expr).EnsureValue() == match {
				e.EvalCodeBlock(c.Body)
				return
			}
		}
	}
	if m.Default != nil {
		e.EvalCodeBlock(m.Default)
	}
}
