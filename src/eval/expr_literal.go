package eval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func (e *Eval) EvalStringLiteral(n *node.StringLiteral) string {
	return n.Value
}

func (e *Eval) EvalFloatLiteral(n *node.FloatLiteral) float64 {
	return n.Value
}

func (e *Eval) EvalIntLiteral(n *node.IntLiteral) int {
	return n.Value
}

func (e *Eval) EvalBoolLiteral(n *node.BoolLiteral) bool {
	return n.Value
}

func (e *Eval) EvalArrayLiteral(n *node.ArrayLiteral) []any {
	var arr []any
	for _, v := range n.Value {
		arr = append(arr, e.EvalExpr(v))
	}
	return arr
}

func (e *Eval) EvalMapLiteral(n *node.MapLiteral) map[any]any {
	var m = make(map[any]any)
	for _, v := range n.Value {
		m[e.EvalExpr(v.Key)] = e.EvalExpr(v.Value)
	}
	return m
}

func (e *Eval) EvalIdentifier(n *node.Identifier) any {
	v, ok := e.objTable.Get(n.Value)
	if ok {
		return v
	}
	panic("No Var Found")
}
