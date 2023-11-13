package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

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
			//e.objTable.Set(baseV.Name.Value, obj)
			//return v
		}
	}
	switch v.(type) {
	case *node.LambdaExpr:
		e.funcTable.Set(baseV.Name.Value, v.(*node.LambdaExpr).ToFunc(baseV.Name.Value))
	default:
		e.objTable.Set(baseV.Name.Value, v)

	}
	e.objTable.Set(baseV.Name.Value, v)
	return v
}
