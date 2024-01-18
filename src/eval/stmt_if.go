package eval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func (e *Eval) EvalIfStmt(n *node.IfStmt) {
	con := e.EvalExpr(n.Condition)
	switch con.(type) {
	case bool:
		if con.(bool) {
			e.EvalCodeBlock(n.IfTrue)
			return
		}
		if n.IfFalse != nil {
			e.EvalCodeBlock(n.IfFalse)
		}
		return
	}
	panic("NOT BOOL")
}
