package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/builtin"
)

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

func (e *Eval) EvalBuiltInCall(fc *node.FuncCall, args []any) any {
	fn := builtin.Match(fc.Caller)
	if fn == nil {
		panic("func not found")
	}
	_ = builtin.Call(fn, args)
	return nil
}
