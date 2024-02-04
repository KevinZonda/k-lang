package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
)

func (e *Eval) EvalCodeBlock(fc *node.CodeBlock) any {
	e.currentToken = fc.GetToken()
	e.frameStart(false)

	_ = e.runAst(fc.Nodes, reserved.Return) // FIXME: return?
	//fe := e.new((tree.Ast)(fc.Nodes))
	//_ = fe.run()
	e.frameEndWithAll()
	return nil
}

func (e *Eval) EvalLoopCodeBlock(fc *node.CodeBlock) any {
	e.currentToken = fc.GetToken()
	e.frameStart(false)

	_ = e.runAst(fc.Nodes, reserved.Return, reserved.Break, reserved.Continue)
	//fe := e.new((tree.Ast)(fc.Nodes))
	//_ = fe.run()
	e.frameEndWithAll()
	return nil
}
