package eval

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"

func (e *Eval) EvalMain() any {
	fn, ok := e.funcTable.Get("main")
	if !ok {
		return nil
	}
	e.frameStart()
	fe := e.new((tree.Ast)(fn.Body.Nodes))
	ret := fe.run()
	e.frameEnd()
	return ret
}
