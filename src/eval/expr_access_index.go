package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
	"slices"
)

func accessArr[T any](arr []T, startIdx int, endIdx int) []T {
	if endIdx > startIdx {
		return arr[startIdx:endIdx]
	}
	s := arr[endIdx:startIdx]
	slices.Reverse(s)
	return s
}

func (e *Eval) EvalIndexExpr(n *node.IndexExpr) any {
	left := e.EvalExpr(n.Left).EnsureValue()
	idx := e.EvalExpr(n.Index).EnsureValue()
	e.currentToken = n.GetToken()
	switch leftT := left.(type) {
	case string:
		arr := []rune(leftT)
		startIdx := asType[int](idx)
		if n.EndIndex == nil {
			if n.Col {
				return string(arr[startIdx:])
			}
			return string(arr[startIdx])
		}
		endIdx := asType[int](e.EvalExpr(n.EndIndex).EnsureValue())
		if endIdx == startIdx {
			return string(arr[startIdx])
		}
		return string(accessArr[rune](arr, startIdx, endIdx))
	case []any:
		//if idx, ok := idx.(int); ok {
		//	if idx < 0 || idx >= len(arr) {
		//	}
		//}
		// FIXME: PANIC INFORMATION
		startIdx := asType[int](idx)
		if n.EndIndex == nil {
			if n.Col {
				return leftT[startIdx:]
			}
			return leftT[startIdx]
		}
		endIdx := asType[int](e.EvalExpr(n.EndIndex).EnsureValue())
		if endIdx == startIdx {
			return leftT[startIdx]
		}
		return accessArr[any](leftT, startIdx, endIdx)
	case map[any]any:
		if n.EndIndex != nil {
			panic("map does not support end index")
		}

		return leftT[idx]
	}
	panic(fmt.Sprint("index expr not impl. -> ", left, reflect.TypeOf(left)))
	return nil
}
