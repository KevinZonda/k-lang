package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
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

	return e.EvalBuiltInCall(fc, args)
}
