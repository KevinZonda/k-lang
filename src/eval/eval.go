package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type Eval struct {
	ast       tree.Ast
	objTable  *obj.TableStack
	funcTable *obj.StackImpl[*node.FuncBlock]
	basePath  string
	loopLvl   int
	opened    map[string]*Eval
}

func (e *Eval) LoadContext(o *Eval) {
	if o == nil {
		e.objTable = &obj.TableStack{}
		e.funcTable = &obj.StackImpl[*node.FuncBlock]{}
		return
	}
	e.objTable = o.objTable
	e.funcTable = o.funcTable
}

func (e *Eval) run() any {
	for _, n := range e.ast {
		switch n.(type) {
		case *node.CodeBlock:
			e.EvalCodeBlock(n.(*node.CodeBlock))
		case node.Expr:
			e.EvalExpr(n.(node.Expr))
		case node.Stmt:
			e.EvalStmt(n.(node.Stmt))
		case *node.FuncBlock:
			fb := n.(*node.FuncBlock)
			e.funcTable.Set(fb.Name.Value, fb)
		case *node.OpenBlock:
			ob := n.(*node.OpenBlock)
			for _, stmt := range ob.Openers {
				e.EvalStmt(stmt)
			}
		default:
			panic("not implemented")
		}
		if e.objTable.HasKeyAtTop(reserved.Return) ||
			e.objTable.HasKeyAtTop(reserved.Break) ||
			e.objTable.HasKeyAtTop(reserved.Continue) {
			break
		}
	}
	val, _ := e.objTable.Get(reserved.Return)
	return val
}

func (e *Eval) It() any {
	v, _ := e.objTable.Get("it")
	return v
}

func (e *Eval) Do() {
	if retV := e.run(); retV != nil {
		fmt.Println("Program returned: ", retV)
	}

	e.EvalMain()
}
