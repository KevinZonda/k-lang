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
	loopLvl   int
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
		default:
			panic("not implemented")
		}
		if e.objTable.HasKeyAtTop(reserved.Return) || e.objTable.HasKeyAtTop(reserved.Break) {
			break
		}
	}
	val, _ := e.objTable.Get("0") // 0 is reserved place for return value
	return val
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
		case *node.FuncBlock:
			fb := n.(*node.FuncBlock)
			e.funcTable.Set(fb.Name.Value, fb)
		default:
			panic("not implemented")
		}
		if retVal, ok := e.objTable.GetAtTop(reserved.Return); ok {
			fmt.Println("Program returned: ", retVal)
		}
	}

	e.EvalMain()
}
