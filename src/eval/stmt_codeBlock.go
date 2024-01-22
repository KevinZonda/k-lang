package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
)

func (e *Eval) EvalCodeBlock(fc *node.CodeBlock) any {
	e.frameStart(false)
	fe := e.new((tree.Ast)(fc.Nodes))
	_ = fe.run()
	e.frameEndWithAll()
	return nil
}
