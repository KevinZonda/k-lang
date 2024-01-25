package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"reflect"
)

func (e *Eval) EvalStmt(n node.Stmt) {
	e.currentToken = n.GetToken()
	switch nT := n.(type) {
	case *node.AssignStmt:
		e.EvalAssignStmt(nT)
	case *node.IfStmt:
		e.EvalIfStmt(nT)
	case *node.FuncCall:
		e.EvalFuncCall(nT)
	case *node.ReturnStmt:
		e.EvalReturnStmt(nT)
	case *node.WhileStyleFor:
		e.EvalWhileForStmt(nT)
	case *node.CStyleFor:
		e.EvalCStyleFrStmt(nT)
	case *node.IterStyleFor:
		e.EvalIterStyleForStmt(nT)
	case *node.BreakStmt:
		e.EvalBreakStmt(nT)
	case *node.ContinueStmt:
		e.EvalContinueStmt(nT)
	case *node.TryCatchStmt:
		e.EvalTryCatchStmt(nT)
	case *node.MatchStmt:
		e.EvalMatchStmt(nT)
	case *node.OpenStmt:
		e.EvalOpenStmt(nT)
	default:
		panic("not implemented for eval stmt type: " + reflect.TypeOf(n).String())
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
