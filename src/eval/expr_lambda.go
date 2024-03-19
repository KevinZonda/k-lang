package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/memory"
)

func (e *Eval) EvalLambdaExpr(n *node.LambdaExpr) *node.LambdaExpr {
	e.currentToken = n.GetToken()
	// closure memory should be ref to current memory section
	if e.memory.Len() <= 1 {
		n.Mem = nil // It's at button, no need to create mem sec
	} else {
		mem := memory.NewMemory()
		for l := e.memory.Len() - 1; l > 0; l-- {
			m := e.memory.Raw()[l]
			_l := m.Copy()
			_l.Protect = false
			mem.Push(_l)
			if m.Protect {
				break
			}
		}
		n.Mem = mem
	}
	return n
}
