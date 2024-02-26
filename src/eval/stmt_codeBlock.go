package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
)

func (e *Eval) EvalCodeBlock(fc *node.CodeBlock) {
	e.currentToken = fc.GetToken()
	e.frameStart(false)

	if e.IsLoopFrame() {
		_ = e.runAst(fc.Nodes, reserved.Return, reserved.Continue, reserved.Break)
	} else {
		_ = e.runAst(fc.Nodes, reserved.Return)
	}
	//fe := e.new((tree.Ast)(fc.Nodes))
	//_ = fe.run()
	e.frameEndWithAll()
}
