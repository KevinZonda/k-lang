package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"path/filepath"
	"reflect"
)

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

func (e *Eval) frameStart() {
	e.objTable.PushEmpty()
	e.funcTable.PushEmpty()
}

func (e *Eval) frameEnd() {
	e.objTable.Pop()
	e.funcTable.Pop()
}

func (e *Eval) frameEndWith(keys ...string) {
	m := e.objTable.Pop()
	e.funcTable.Pop()
	for _, key := range keys {
		if v, ok := m[key]; ok {
			switch v.(type) {
			case *node.LambdaExpr:
				e.funcTable.Set(key, v.(*node.LambdaExpr).ToFunc(key))
			default:
				e.objTable.Set(key, v)
			}
		}
	}
}

func (e *Eval) frameEndWithAll() any {
	e.frameEndWith(reserved.Return, reserved.Break, reserved.Continue)
	retV, _ := e.objTable.Peek()[reserved.Return]
	return retV
}

func (e *Eval) evalValWithIdx(idxs []node.Expr, root *any) *any {
	val := root
	for _, idxExpr := range idxs {
		idx := e.EvalExpr(idxExpr)
		switch reflect.TypeOf(*val).Kind() {
		case reflect.Slice, reflect.Array:
			*val = (*val).([]any)[idx.(int)]
		case reflect.Map:
			//if reflect.TypeOf(*val).Key() != reflect.TypeOf(idx) {
			//	panic("Invalid key type")
			//}
			*val = (*val).(map[any]any)[idx]
		}
	}
	return val
}

func (e *Eval) ObjTable() *obj.TableStack {
	return e.objTable
}

func (e *Eval) FuncTable() *obj.StackImpl[*node.FuncBlock] {
	return e.funcTable
}
