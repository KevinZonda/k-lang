package eval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func (e *Eval) EvalIfStmt(n *node.IfStmt) {
	e.currentToken = n.GetToken()
	con := e.EvalExpr(n.Condition)
	if conV, ok := con.(bool); ok {
		if conV {
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
