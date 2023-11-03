package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

func (e *Eval) EvalStmt(n node.Stmt) any {
	switch n.(type) {
	case *node.AssignStmt:
		return e.EvalAssignStmt(n.(*node.AssignStmt))
	case *node.IfStmt:
		return e.EvalIfStmt(n.(*node.IfStmt))
	case *node.FuncCall:
		return e.EvalFuncCall(n.(*node.FuncCall))
	case *node.ReturnStmt:
		return e.EvalReturnStmt(n.(*node.ReturnStmt))
	case *node.WhileStyleFor:
		return e.EvalWhileForStmt(n.(*node.WhileStyleFor))
	case *node.BreakStmt:
		return e.EvalBreakStmt(n.(*node.BreakStmt))
	}
	panic("not implemented")
}

func (e *Eval) EvalBreakStmt(n *node.BreakStmt) any {
	if e.loopLvl <= 0 {
		return nil
	}
	e.objTable.SetAtTop("1", true)
	return nil
}

func (e *Eval) EvalReturnStmt(n *node.ReturnStmt) any {
	e.objTable.SetAtTop("0", nil)
	if n.Value != nil {
		e.objTable.SetAtTop("0", e.EvalExpr(n.Value))
	}
	return nil
}
