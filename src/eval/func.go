package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/builtin"
)

func (e *Eval) EvalBuiltInCall(fc *node.FuncCall, args []any) any {
	fn := builtin.Match(fc.Caller)
	if fn == nil {
		panic("func not found")
	}
	_ = builtin.Call(fn, args)
	return nil
}

func (e *Eval) EvalFuncCall(fc *node.FuncCall) any {
	exprArgs := fc.Args
	var args []any
	for _, expr := range exprArgs {
		args = append(args, e.EvalExpr(expr))
	}

	funcName := fc.Caller.Value[0].Name.Value
	fn, ok := e.funcTable.Get(funcName)
	if !ok {
		return e.EvalBuiltInCall(fc, args)
	}

	e.frameStart()
	for i, funcArg := range fn.Args {
		e.objTable.Set(funcArg.Name.Value, args[i])
	}
	fe := e.new((tree.Ast)(fn.Body.Nodes))
	_ = fe.run()
	retV, retOk := fe.objTable.GetAtTop("0")
	e.frameEnd()
	if retOk {
		return retV
	}
	return nil
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

func (e *Eval) EvalCodeBlock(fc *node.CodeBlock) any {
	e.frameStart()
	fe := e.new((tree.Ast)(fc.Nodes))
	_ = fe.run()
	e.frameEndWithRetAndBreak()
	return nil
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
