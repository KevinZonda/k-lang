package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
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
			if e.objTable.HasKeyAtTop("0") || e.objTable.HasKeyAtTop("1") {
				break
			}
		case *node.FuncBlock:
			fb := n.(*node.FuncBlock)
			e.funcTable.Set(fb.Name.Value, fb)
		default:
			panic("not implemented")
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
			v, ok := e.objTable.GetAtTop("0")
			if ok {
				fmt.Println("Program returned: ", v)
				// Program Exit
			}
		case *node.FuncBlock:
			fb := n.(*node.FuncBlock)
			e.funcTable.Set(fb.Name.Value, fb)
		default:
			panic("not implemented")
		}
	}

	e.EvalMain()

}

func New(ast tree.Ast) *Eval {
	return &Eval{
		ast:       ast,
		objTable:  &obj.TableStack{},
		funcTable: &obj.StackImpl[*node.FuncBlock]{},
	}
}

func (e *Eval) new(ast tree.Ast) *Eval {
	return &Eval{
		ast:       ast,
		objTable:  e.objTable,
		funcTable: e.funcTable,
		loopLvl:   e.loopLvl,
	}
}
