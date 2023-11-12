package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
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
	case *node.CStyleFor:
		return e.EvalCStyleFrStmt(n.(*node.CStyleFor))
	case *node.BreakStmt:
		return e.EvalBreakStmt(n.(*node.BreakStmt))
	case *node.ContinueStmt:
		return e.EvalContinueStmt(n.(*node.ContinueStmt))
	case *node.MatchStmt:
		return e.EvalMatchStmt(n.(*node.MatchStmt))
	}
	panic("not implemented")
}

func (e *Eval) EvalBreakStmt(n *node.BreakStmt) any {
	if e.loopLvl <= 0 {
		return nil
	}
	e.objTable.SetAtTop(reserved.Break, true)
	return nil
}

func (e *Eval) EvalContinueStmt(n *node.ContinueStmt) any {
	if e.loopLvl <= 0 {
		return nil
	}
	e.objTable.SetAtTop(reserved.Continue, true)
	return nil
}

func (e *Eval) EvalReturnStmt(n *node.ReturnStmt) any {
	e.objTable.SetAtTop(reserved.Return, nil)
	if n.Value != nil {
		e.objTable.SetAtTop(reserved.Return, e.EvalExpr(n.Value))
	}
	return nil
}
