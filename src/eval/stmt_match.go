package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

func (e *Eval) EvalMatchStmt(m *node.MatchStmt) any {
	match := e.EvalExpr(m.Match)
	for _, c := range m.Cases {
		if e.EvalExpr(c.Expr) == match {
			return e.EvalCodeBlock(c.Body)
		}
	}
	if m.Default != nil {
		return e.EvalCodeBlock(m.Default)
	}
	return nil
}
