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
	case *node.DeclareStmt:
		e.EvalDeclareStmt(nT)
	default:
		panic("not implemented for eval stmt type: " + reflect.TypeOf(n).String())
	}
}

func (e *Eval) EvalBreakStmt(n *node.BreakStmt) {
	e.currentToken = n.GetToken()
	if e.loopLvl <= 0 {
		return
	}
	e.memory.Top().SetValue(reserved.Break, true)
	return
}

func (e *Eval) EvalContinueStmt(n *node.ContinueStmt) {
	e.currentToken = n.GetToken()
	if e.loopLvl <= 0 {
		return
	}
	e.memory.Top().SetValue(reserved.Continue, true)
	return
}

func (e *Eval) EvalReturnStmt(n *node.ReturnStmt) {
	e.currentToken = n.GetToken()
	if n.Value == nil {
		e.memory.Top().SetValue(reserved.Return, nil)
		return
	}
	if len(n.Value) == 1 {
		e.memory.Top().SetValue(reserved.Return, e.EvalExpr(n.Value[0]).EnsureValue())
		return
	}
	e.memory.Top().SetValue(reserved.Return, e.evalExprs(n.Value...))
}

func (e *Eval) evalReturnVal() (vals []any, hasRet bool) {
	v, hasRet := e.memory.Top().GetValue(reserved.Return)
	if !hasRet {
		return nil, false
	}
	if v == nil {
		return nil, true
	}
	switch t := v.(type) {
	case []any:
		return t, true
	default:
		return []any{t}, true
	}
}
