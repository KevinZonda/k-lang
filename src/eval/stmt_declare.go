package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) EvalDeclareStmt(n *node.DeclareStmt) {
	val := e.getZeroValue(n.Type)
	for _, name := range n.Names {
		e.memory.Set(name, obj.Construct(val).WithType(n.Type))
	}
}
