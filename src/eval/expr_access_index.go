package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"log"
	"reflect"
)

func (e *Eval) EvalIndexExpr(n *node.IndexExpr) any {
	left := e.EvalExpr(n.Left)
	idx := e.EvalExpr(n.Index)
	e.currentToken = n.GetToken()
	if arr, ok := left.([]any); ok {
		//if idx, ok := idx.(int); ok {
		//	if idx < 0 || idx >= len(arr) {
		//	}
		//}
		// FIXME: PANIC INFORMATION
		return arr[idx.(int)]
	}
	if m, ok := left.(map[any]any); ok {
		return m[idx]
	}
	log.Println("Left: ", left, reflect.TypeOf(left))
	panic("not implemented")
	return nil
}
