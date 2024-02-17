package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
)

func (e *Eval) EvalIndexExpr(n *node.IndexExpr) any {
	left := e.EvalExpr(n.Left).EnsureValue()
	idx := e.EvalExpr(n.Index).EnsureValue()
	e.currentToken = n.GetToken()
	switch leftT := left.(type) {
	case string:
		return string([]rune(leftT)[asType[int](idx)])
	case []any:
		//if idx, ok := idx.(int); ok {
		//	if idx < 0 || idx >= len(arr) {
		//	}
		//}
		// FIXME: PANIC INFORMATION
		return leftT[asType[int](idx)]
	case map[any]any:
		return leftT[idx]
	}
	panic(fmt.Sprint("index expr not impl. -> ", left, reflect.TypeOf(left)))
	return nil
}
