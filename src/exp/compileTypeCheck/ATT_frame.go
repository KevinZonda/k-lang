package compileTypeCheck

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/compileTypeCheck/tmem"

func (e *ATT) frameStart(protect bool) *TMem.TMem {
	return e.m.PushEmpty(protect)
}

func (e *ATT) frameTopProtection(protect bool) {
	if top := e.m.Top(); top != nil {
		top.Protect = protect
	}
}
func (e *ATT) frameEnd() {
	e.m.Pop()
}
