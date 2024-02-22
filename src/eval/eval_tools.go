package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/memory"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
)

func (e *Eval) frameStart(protect bool) *memory.Layer {
	return e.memory.PushEmpty(protect)
}

func (e *Eval) frameTopProtection(protect bool) {
	if top := e.memory.Top(); top != nil {
		top.Protect = protect
	}
}
func (e *Eval) frameEnd() {
	e.memory.Pop()
}

func (e *Eval) frameEndWith(keys ...string) {
	m := e.memory.Pop()
	for _, key := range keys {
		if v, ok := m.Get(key); ok {
			e.memory.Set(key, v)
		}
	}
}

func (e *Eval) frameEndWithAll() any {
	e.frameEndWith(reserved.Return, reserved.Break, reserved.Continue)
	peek := e.memory.Top()
	if peek.Empty() {
		return nil
	}
	retV, _ := e.memory.Top().Get(reserved.Return)
	return retV
}
