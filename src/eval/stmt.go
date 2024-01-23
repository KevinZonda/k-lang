package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
)

func (e *Eval) EvalStmt(n node.Stmt) {
	e.currentToken = n.GetToken()
	switch n.(type) {
	case *node.AssignStmt:
		e.EvalAssignStmt(n.(*node.AssignStmt))
	case *node.IfStmt:
		e.EvalIfStmt(n.(*node.IfStmt))
	case *node.FuncCall:
		e.EvalFuncCall(n.(*node.FuncCall))
	case *node.ReturnStmt:
		e.EvalReturnStmt(n.(*node.ReturnStmt))
	case *node.WhileStyleFor:
		e.EvalWhileForStmt(n.(*node.WhileStyleFor))
	case *node.CStyleFor:
		e.EvalCStyleFrStmt(n.(*node.CStyleFor))
	case *node.IterStyleFor:
		e.EvalIterStyleForStmt(n.(*node.IterStyleFor))
	case *node.BreakStmt:
		e.EvalBreakStmt(n.(*node.BreakStmt))
	case *node.ContinueStmt:
		e.EvalContinueStmt(n.(*node.ContinueStmt))
	case *node.TryCatchStmt:
		e.EvalTryCatchStmt(n.(*node.TryCatchStmt))
	case *node.MatchStmt:
		e.EvalMatchStmt(n.(*node.MatchStmt))
	case *node.OpenStmt:
		e.EvalOpenStmt(n.(*node.OpenStmt))
	default:
		panic("not implemented")
	}
}

func (e *Eval) EvalBreakStmt(n *node.BreakStmt) {
	e.currentToken = n.GetToken()
	if e.loopLvl <= 0 {
		return
	}
	e.objTable.SetAtTop(reserved.Break, true)
	return
}

func (e *Eval) EvalContinueStmt(n *node.ContinueStmt) {
	e.currentToken = n.GetToken()
	if e.loopLvl <= 0 {
		return
	}
	e.objTable.SetAtTop(reserved.Continue, true)
	return
}

func (e *Eval) EvalReturnStmt(n *node.ReturnStmt) {
	e.currentToken = n.GetToken()
	e.objTable.SetAtTop(reserved.Return, nil)
	if n.Value != nil {
		e.objTable.SetAtTop(reserved.Return, e.EvalExpr(n.Value))
	}
	return
}
