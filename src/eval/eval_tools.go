package eval

import (
	"reflect"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
)

func (e *Eval) frameStart(protect bool) {
	e.objTable.PushEmpty(protect)
}

func (e *Eval) frameEnd() {
	e.objTable.Pop()
}

func (e *Eval) frameEndWith(keys ...string) {
	m := e.objTable.Pop()
	for _, key := range keys {
		if v, ok := m.Get(key); ok {
			e.objTable.Set(key, v)
		}
	}
}

func (e *Eval) frameEndWithAll() any {
	e.frameEndWith(reserved.Return, reserved.Break, reserved.Continue)
	peek := e.objTable.Peek()
	if peek.Empty() {
		return nil
	}
	retV, _ := e.objTable.Peek().Get(reserved.Return)
	return retV
}

func (e *Eval) evalValWithIdx(idxs []node.Expr, root *any) *any {
	val := root
	for _, idxExpr := range idxs {
		idx := e.EvalExpr(idxExpr)
		switch reflect.TypeOf(*val).Kind() {
		case reflect.Slice, reflect.Array:
			*val = (*val).([]any)[idx.(int)]
		case reflect.Map:
			//if reflect.TypeOf(*val).Key() != reflect.TypeOf(idx) {
			//	panic("Invalid key type")
			//}
			*val = (*val).(map[any]any)[idx]
		}
	}
	return val
}

func (e *Eval) ObjTable() *TableStack {
	return e.objTable
}
