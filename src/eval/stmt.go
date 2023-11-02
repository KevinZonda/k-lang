package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
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
		return nil
	}
	panic("not implemented")
}

func (e *Eval) EvalWhileForStmt(n *node.WhileStyleFor) any {
	var r any
	for {
		if n.ConditionExpr != nil {
			rst := e.EvalExpr(n.ConditionExpr).(bool)
			if !rst {
				return nil
			}
		}
		r = e.EvalLoopCodeBlock(n.Body)
	}
	return r
}

func (e *Eval) EvalLoopCodeBlock(fc *node.CodeBlock) (val any) {
	e.objTable.PushEmpty()
	e.funcTable.PushEmpty()
	fe := new((tree.Ast)(fc.Nodes), e.objTable, e.funcTable)
	_ = fe.run()
	val, ok := fe.objTable.GetAtTop("0")
	if !ok {
		val = nil
	}
	e.objTable.Pop()
	e.funcTable.Pop()
	return val
}

func (e *Eval) EvalReturnStmt(n *node.ReturnStmt) any {
	e.objTable.SetAtTop("0", nil)
	if n.Value != nil {
		e.objTable.SetAtTop("0", e.EvalExpr(n.Value))
	}
	return nil
}

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

func (e *Eval) EvalAssignStmt(n *node.AssignStmt) any {
	v := e.EvalExpr(n.Value)
	// TODO: arr
	baseV := n.Var.Value[len(n.Var.Value)-1]
	obj, ok := e.objTable.Get(baseV.Name.Value)
	if ok && len(baseV.Index) != 0 {
		switch obj.(type) {
		case []any:
			bIdx := baseV.Index[:len(baseV.Index)-1]
			if tgt := e.evalValWithIdx(bIdx, &obj); tgt != nil {
				switch (*tgt).(type) {
				case []any:
					((*tgt).([]any))[e.EvalExpr(baseV.Index[len(baseV.Index)-1]).(int)] = v
				}
			}
			e.objTable.Set(baseV.Name.Value, obj)
			return v
		}
	}
	e.objTable.Set(baseV.Name.Value, v)
	return v
}

func (e *Eval) evalValWithIdx(idxs []node.Expr, root *any) *any {
	val := root
	for _, idxExpr := range idxs {
		idx := e.EvalExpr(idxExpr)
		switch idx.(type) {
		case int:
			*val = (*val).([]any)[idx.(int)]
		}
	}
	return val
}
