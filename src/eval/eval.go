package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type Eval struct {
	ast      tree.Ast
	objTable *obj.TableStack
}

func (e *Eval) It() any {
	v, _ := e.objTable.Get("it")
	return v
}

func (e *Eval) Do() {
	for _, n := range e.ast {
		switch n.(type) {
		case node.Expr:
			e.objTable.Set("it", e.EvalExpr(n.(node.Expr)))
		case node.Stmt:
			e.EvalStmt(n.(node.Stmt))
		default:
			panic("not implemented")
		}
	}
}

func New(ast tree.Ast) *Eval {
	return &Eval{
		ast:      ast,
		objTable: &obj.TableStack{},
	}
}
