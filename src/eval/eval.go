package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"path/filepath"
)

type Eval struct {
	ast       tree.Ast
	objTable  *obj.TableStack
	funcTable *obj.StackImpl[*node.FuncBlock]
	basePath  string
	loopLvl   int
	opened    map[string]*Eval
}

var openedFiles map[string]*Eval

func New(ast tree.Ast, inputFile string) *Eval {
	if openedFiles == nil {
		openedFiles = map[string]*Eval{}
	}
	path := filepath.Dir(inputFile)
	return &Eval{
		ast:       ast,
		objTable:  obj.NewObjectTable(),
		funcTable: &obj.StackImpl[*node.FuncBlock]{},
		basePath:  path,
		opened:    make(map[string]*Eval),
	}
}

func (e *Eval) new(ast tree.Ast) *Eval {
	return &Eval{
		ast:       ast,
		objTable:  e.objTable,
		funcTable: e.funcTable,
		loopLvl:   e.loopLvl,
		basePath:  e.basePath,
		opened:    e.opened,
	}
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

func (e *Eval) runWithBreak(breaks ...string) any {
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
		for _, key := range breaks {
			if e.objTable.HasKeyAtTop(key) {
				break
			}
		}
	}
	val, _ := e.objTable.Get(reserved.Return)
	return val
}

func (e *Eval) run() any {
	return e.runWithBreak(reserved.Return, reserved.Break, reserved.Continue)
}

func (e *Eval) Do() {
	if retV := e.runWithBreak(reserved.Return); retV != nil {
		fmt.Println("[Interpreter] Program returned: ", retV)
		return
	}

	e.EvalMain()
}

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
