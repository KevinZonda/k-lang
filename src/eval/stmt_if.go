package eval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func (e *Eval) EvalIfStmt(n *node.IfStmt) any {
	con := e.EvalExpr(n.Condition)
	switch con.(type) {
	case bool:
		if con.(bool) {
			return e.EvalCodeBlock(n.IfTrue)
		}
		if n.IfFalse != nil {
			return e.EvalCodeBlock(n.IfFalse)
		}
		return nil
	}
	panic("NOT BOOL")
}
