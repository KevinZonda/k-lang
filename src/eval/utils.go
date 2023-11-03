package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

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
			e.objTable.SetAtTop(key, v)
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
		switch idx.(type) {
		case int:
			*val = (*val).([]any)[idx.(int)]
		}
	}
	return val
}
