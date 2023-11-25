package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) EvalFuncCall(fc *node.FuncCall) any {
	exprArgs := fc.Args
	var args []any
	for _, expr := range exprArgs {
		args = append(args, e.EvalExpr(expr))
	}

	funcName := fc.Caller.Value
	fx, ok := e.objTable.Get(funcName)
	if !ok || !(fx.Is(obj.Func, obj.Lambda)) {
		return e.EvalBuiltInCall(fc, args)
	}
	var fn *node.FuncBlock
	if fx.Is(obj.Lambda) {
		fn = fx.ToLambda().ToFunc(funcName)
	} else {
		fn = fx.ToFunc()
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
	fn := builtin.Match(fc.Caller.Value)
	if fn == nil {
		panic("func not found")
	}
	var ret []any
	xs := builtin.Call(fn, args)
	for _, x := range xs {
		ret = append(ret, x)
	}
	// TODO: return multiple values
	if len(ret) > 0 {
		return ret[0]
	}
	return nil
}
