package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) evalExprs(exprs ...node.Expr) []any {
	var ret []any
	for _, expr := range exprs {
		ret = append(ret, e.EvalExpr(expr))
	}
	return ret
}

func (e *Eval) EvalFuncCall(fc *node.FuncCall) any {
	e.currentToken = fc.GetToken()
	// Eval Args

	// Find Func
	funcName := fc.Caller.Value
	fx, ok := e.objTable.Get(funcName)
	if !ok || !(fx.Is(obj.Func, obj.Lambda)) {
		return e.EvalBuiltInCall(fc, e.evalExprs(fc.Args...))
	}

	var fn *node.FuncBlock
	if fx.Is(obj.Lambda) {
		fn = fx.ToLambda().ToFunc(funcName)
	} else {
		fn = fx.ToFunc()
	}

	return e.EvalFuncBlock(fn, fc.Args, nil)
	//e.frameStart()
	//for i, funcArg := range fn.Args {
	//	e.objTable.Set(funcArg.Name.Value, args[i])
	//}
	//fe := e.new((tree.Ast)(fn.Body.Nodes))
	//_ = fe.run()
	//retV, retOk := fe.objTable.GetAtTop("0")
	//e.frameEnd()
	//if retOk {
	//	return retV
	//}
	//return nil
}

func (e *Eval) EvalFuncBlock(fn *node.FuncBlock, args []node.Expr, onAfterFrameStart func()) any {
	e.currentToken = fn.GetToken()
	e.frameStart(true)
	if onAfterFrameStart != nil {
		onAfterFrameStart()
	}
	for i, funcArg := range fn.Args {
		var v any = nil
		if funcArg.Ref {
			v = e.evalExpr(args[i], true)
		} else {
			v = clone(e.EvalExpr(args[i]))
		}
		e.objTable.SetAtTop(funcArg.Name.Value, v)
	}

	result := e.runAst((tree.Ast)(fn.Body.Nodes), reserved.Return)

	//fe := e.new((tree.Ast)(fn.Body.Nodes))
	//_ = fe.run()
	// retV, retOk := fe.objTable.GetAtTop("0")
	e.frameEnd()
	if result.hasRet {
		return result.retV.Val
	}
	return nil
}

func (e *Eval) Mem() {
	e.objTable.Println(e.builtin.StdOut, e.PtrAddr())
}

func (e *Eval) EvalBuiltInCall(fc *node.FuncCall, args []any) any {
	e.currentToken = fc.GetToken()
	if fc.Caller.Value == "MEM" {
		e.Mem()
		return nil
	}
	fn := e.builtin.Match(fc.Caller.Value)
	if fn == nil {
		panic("func not found")
	}
	var ret []any
	xs := e.builtin.Call(fn, args)
	for _, x := range xs {
		ret = append(ret, x)
	}
	// TODO: return multiple values
	if len(ret) > 0 {
		return ret[0]
	}
	return nil
}
