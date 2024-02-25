package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) evalExprs(exprs ...node.Expr) []any {
	var ret []any
	for _, expr := range exprs {
		ret = append(ret, e.EvalExpr(expr).EnsureValue())
	}
	return ret
}

func (e *Eval) EvalFuncCall(fc *node.FuncCall) ExprResult {
	e.currentToken = fc.GetToken()
	// Eval Args

	// Find Func
	funcName := fc.Caller.Value
	// TODO: int() float()
	fx, ok := e.memory.Get(funcName)
	if !ok || !(fx.Is(obj.Func, obj.Lambda)) {
		if fx.Is(obj.StructDef) {
			t := &node.Type{
				Name: funcName,
			}
			v := e.EvalExpr(fc.Args[0]).EnsureValue()
			e.TypeCheckOrPanic(t, v)
			retV := e.NormaliseWithType(t, v)
			return exprVal(retV)
		}
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
	//	e.memory.Set(funcArg.Name.Value, args[i])
	//}
	//fe := e.new((tree.Ast)(fn.Body.Nodes))
	//_ = fe.run()
	//retV, retOk := fe.memory.GetAtTop("0")
	//e.frameEnd()
	//if retOk {
	//	return retV
	//}
	//return nil
}

func (e *Eval) EvalFuncBlock(fn *node.FuncBlock, args []node.Expr, onNewFrame func()) ExprResult {
	e.currentToken = fn.GetToken()
	topFrame := e.frameStart(false)
	if onNewFrame != nil {
		onNewFrame()
	}
	for i, funcArg := range fn.Args {
		var v any = nil
		if funcArg.Ref {
			v = e.evalExpr(args[i], true).EnsureValue()
			if vT, ok := v.(*obj.Object); ok {
				v = vT.CreateRef()
			}
		} else {
			v = clone(e.EvalExpr(args[i]).EnsureValue())
		}
		e.TypeCheckOrPanic(funcArg.Type, v)
		v = e.NormaliseWithType(funcArg.Type, v)
		e.memory.Top().SetValue(funcArg.Name.Value, v)
	}

	topFrame.Protect = true
	result := e.runAst((tree.Ast)(fn.Body.Nodes), reserved.Return)

	//fe := e.new((tree.Ast)(fn.Body.Nodes))
	//_ = fe.run()
	// retV, retOk := fe.memory.GetAtTop("0")
	e.frameEnd()
	if fn.RetType != nil {
		isVoid := fn.RetType[0].IsPlainType(node.TypeVoid)
		if isVoid {
			if result.HasReturn {
				panic("A void func cannot have return value.")
			}
			goto end
		}

		if !isVoid && !result.HasReturn {
			panic("A non-void func cannot have no return value.")
		}
		switch len(fn.RetType) {
		case 0:
			goto end
		case 1:
			e.TypeCheckOrPanic(fn.RetType[0], result.ReturnValue)
		default:
			retV, ok := result.ReturnValue.([]any)
			if !ok || len(retV) != len(fn.RetType) {
				panic("return value should has the same length as return type")
			}
			for idx, retT := range fn.RetType {
				e.TypeCheckOrPanic(retT, retV[idx])
				retV[idx] = e.NormaliseWithType(retT, retV[idx])
			}
		}

	}
end:
	return ExprResult{HasValue: result.HasReturn, Value: result.ReturnValue}
}

func (e *Eval) Mem() {
	e.memory.Println(e.std.StdOut, e.PtrAddr())
}

func (e *Eval) MemStr() string {
	return e.memory.ToString(e.PtrAddr())
}

func (e *Eval) EvalBuiltInCall(fc *node.FuncCall, args []any) ExprResult {
	e.currentToken = fc.GetToken()
	if fc.Caller.Value == "MEM" {
		e.Mem()
		return exprNoVal()
	}
	v := builtin.Set.Call(e.std, fc.Caller.Value, args)
	if v == nil {
		return exprNoVal()
	}
	return exprVal(v)
}
