package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
)

func (e *Eval) frameStart() {
	e.objTable.PushEmpty()
}

func (e *Eval) frameEnd() {
	e.objTable.Pop()
}

func (e *Eval) frameEndWith(keys ...string) {
	m := e.objTable.Pop()
	for _, key := range keys {
		if v, ok := m[key]; ok {
			e.objTable.Set(key, v)
		}
	}
}

func (e *Eval) frameEndWithAll() any {
	e.frameEndWith(reserved.Return, reserved.Break, reserved.Continue)
	retV, _ := e.objTable.Peek()[reserved.Return]
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

func (e *Eval) ObjTable() *obj.TableStack {
	return e.objTable
}
