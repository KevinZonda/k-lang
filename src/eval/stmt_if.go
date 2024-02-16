package eval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func (e *Eval) EvalIfStmt(n *node.IfStmt) {
	e.currentToken = n.GetToken()
	con := e.EvalExpr(n.Condition).EnsureValue()
	if conV, ok := con.(bool); ok {
		if conV {
			e.EvalCodeBlock(n.IfTrue)
			return
		}
		if n.IfFalseIfStmt != nil {
			e.EvalIfStmt(n.IfFalseIfStmt)
			return
		}
		if n.IfFalseCodeBlock != nil {
			e.EvalCodeBlock(n.IfFalseCodeBlock)
		}
		return
	}
	panic("NOT BOOL")
}
