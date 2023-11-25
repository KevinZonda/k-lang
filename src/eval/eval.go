package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/objType"
	"path/filepath"
	"reflect"
)

type Eval struct {
	ast      tree.Ast
	objTable *TableStack
	basePath string
	loopLvl  int
}

var openedFiles map[string]*Eval

func ResetGlobal() {
	openedFiles = make(map[string]*Eval)
}

func New(ast tree.Ast, inputFile string) *Eval {
	if openedFiles == nil {
		openedFiles = map[string]*Eval{}
	}
	path := filepath.Dir(inputFile)
	return &Eval{
		ast:      ast,
		objTable: NewObjectTable(),
		basePath: path,
	}
}

func (e *Eval) new(ast tree.Ast) *Eval {
	return &Eval{
		ast:      ast,
		objTable: e.objTable,
		loopLvl:  e.loopLvl,
		basePath: e.basePath,
	}
}

func (e *Eval) LoadContext(o *Eval) {
	if o == nil {
		e.objTable = &TableStack{}
		return
	}
	e.objTable = o.objTable
}

func (e *Eval) runWithBreak(breaks ...string) (retV *objType.Object, hasRet bool) {
	for _, n := range e.ast {
		for _, key := range breaks {
			if e.objTable.HasKeyAtTop(key) {
				goto end
			}
		}
		switch n.(type) {
		case *node.CodeBlock:
			e.EvalCodeBlock(n.(*node.CodeBlock))
		case node.Expr:
			e.EvalExpr(n.(node.Expr))
		case node.Stmt:
			e.EvalStmt(n.(node.Stmt))
		case *node.FuncBlock:
			fb := n.(*node.FuncBlock)
			e.objTable.Set(fb.Name.Value, fb)
		case *node.StructBlock:
			structBlock := n.(*node.StructBlock)
			e.objTable.Set(structBlock.Name, structBlock)
		case *node.OpenBlock:
			ob := n.(*node.OpenBlock)
			for _, stmt := range ob.Openers {
				e.EvalStmt(stmt)
			}
		default:
			panic("not implemented")
		}
	}
end:
	retV, hasRet = e.objTable.Get(reserved.Return)
	return
}

func (e *Eval) run() any {
	retV, hasRet := e.runWithBreak(reserved.Return, reserved.Break, reserved.Continue)
	if hasRet {
		return retV.Val
	}
	return nil
}

func (e *Eval) Do() {
	if retV, hasRetV := e.runWithBreak(reserved.Return); hasRetV {
		fmt.Println(retV == nil, reflect.TypeOf(retV))
		fmt.Println("[Interpreter] Program returned: ", retV.Val)
		return
	}

	e.EvalMain()
}

func (e *Eval) EvalMain() any {
	fn, ok := e.objTable.Get("main")
	if !ok || !fn.Is(objType.Func) {
		return nil
	}
	fx := fn.ToFunc()
	e.frameStart()
	fe := e.new((tree.Ast)(fx.Body.Nodes))
	ret := fe.run()
	e.frameEnd()
	return ret
}
