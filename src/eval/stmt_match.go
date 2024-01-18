package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

func (e *Eval) EvalMatchStmt(m *node.MatchStmt) {
	match := e.EvalExpr(m.Match)
	for _, c := range m.Cases {
		if e.EvalExpr(c.Expr) == match {
			e.EvalCodeBlock(c.Body)
			return
		}
	}
	if m.Default != nil {
		e.EvalCodeBlock(m.Default)
	}
}
