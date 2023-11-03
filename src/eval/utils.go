package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
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

func (e *Eval) frameEndWithRetAndBreak() {
	retV, hasVal := e.objTable.GetAtTop("0")
	isBreak := e.objTable.HasKeyAtTop("1")
	e.frameEnd()
	if hasVal {
		e.objTable.SetAtTop("0", retV)
	}
	if isBreak {
		e.objTable.SetAtTop("1", nil)
	}
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
