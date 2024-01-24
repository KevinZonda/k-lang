package eval

import (
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

func (e *Eval) ObjTable() *TableStack {
	return e.objTable
}
