package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/memory"
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

func (e *Eval) evalExprsRef(exprs ...node.Expr) []any {
	var ret []any
	for _, expr := range exprs {
		ret = append(ret, e.evalExpr(expr, true).EnsureValue())
	}
	return ret
}

func (e *Eval) funcCallFxPushMem(m any) *EvalFuncBlockEvent {
	if m == nil {
		return nil
	}
	mem := m.(*memory.Layer)
	return &EvalFuncBlockEvent{
		OnNewFramePushed: func() {
			e.memory.Push(mem)
		},
		OnRelease: func() {
			e.memory.Pop()
		},
	}
}

func (e *Eval) EvalFuncCall(fc *node.FuncCall) ExprResult {
	e.currentToken = fc.GetToken()
	// Eval Args

	if fc.CallExpr != nil {
		var fn *node.FuncBlock
		v := e.EvalExpr(fc.CallExpr).EnsureValue()
		var evt *EvalFuncBlockEvent
		switch vT := v.(type) {
		case *node.LambdaExpr:
			fn = &node.FuncBlock{
				Args:    vT.Args,
				Body:    vT.Body,
				RetType: vT.RetType,
			}
			evt = e.funcCallFxPushMem(vT.Mem)
		case *node.FuncBlock:
			fn = vT
		default:
			panic(fmt.Sprintf("Invalid CallExpr: %v", v))
		}
		return e.EvalFuncBlock(fn, fc.Args, evt)
	}
	// Find Func
	funcName := fc.Caller.Value
	fx, ok := e.memory.Get(funcName)
	if !ok || !(fx.Is(obj.Func, obj.Lambda)) {
		if funcName == "visualize" || funcName == "visualise" {
			args := e.evalExprsRef(fc.Args...)
			switch len(args) {
			case 0:
				return exprVal("")
			case 1:
				return exprVal(obj.TreeAnyT("", args[0], true))
			default:
				rst := ""
				for i, arg := range args {
					rst += fmt.Sprintf("arg %d: %s\n", i, obj.TreeAnyT("", arg, true))
				}
				return exprVal(rst)
			}
		}

		if fx.Is(obj.StructDef) {
			t := &node.Type{
				Name: funcName,
			}
			v := e.EvalExpr(fc.Args[0]).EnsureValue()
			retV := e.TypeCast(t, v)
			return exprVal(retV)
		}

		// Internal Type Casting
		switch funcName {
		case node.TypeInt, node.TypeNum, node.TypeString, node.TypeBool:
			t := node.NewType().WithName(funcName)
			v := e.EvalExpr(fc.Args[0]).EnsureValue()
			retV := e.TypeCast(t, v)
			return exprVal(retV)
		}

		return e.EvalBuiltInCall(fc, e.evalExprs(fc.Args...))
	}

	var fn *node.FuncBlock
	var evt *EvalFuncBlockEvent
	if fx.Is(obj.Lambda) {
		lambda := fx.ToLambda()
		fn = lambda.ToFunc(funcName)
		evt = e.funcCallFxPushMem(lambda.Mem)
	} else {
		fn = fx.ToFunc()
	}

	return e.EvalFuncBlock(fn, fc.Args, evt)
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

type EvalFuncBlockEvent struct {
	OnNewFrameCreated func(top *memory.Layer)
	OnNewFramePushed  func()
	OnRelease         func()
}

func (e *Eval) EvalFuncBlock(fn *node.FuncBlock, args []node.Expr, evt *EvalFuncBlockEvent) ExprResult {
	e.currentToken = fn.GetToken()
	if fn.BinaryFx != nil {
		var argsV []any
		for _, arg := range args {
			argsV = append(argsV, e.EvalExpr(arg).EnsureValue())
		}
		vals := fn.EvalBinary(argsV)
		switch len(vals) {
		case 0:
			return exprNoVal()
		case 1:
			return exprVal(vals[0])
		default:
			return exprVal(vals)
		}
	}

	topFrame := memory.NewLayer(false)
	if evt != nil && evt.OnNewFrameCreated != nil {
		evt.OnNewFrameCreated(topFrame)
	}
	for i, funcArg := range fn.Args {
		v := e.evalExprOrRef(args[i], funcArg.Ref)
		v = e.TypeCast(funcArg.Type, v)
		topFrame.SetValue(funcArg.Name.Value, v)
	}
	e.frameStart(true)
	defer e.frameEnd()
	if evt != nil && evt.OnNewFramePushed != nil {
		evt.OnNewFramePushed()
	}
	e.memory.Push(topFrame)
	defer e.frameEnd()

	result := e.runAst((tree.Ast)(fn.Body.Nodes), reserved.Return)

	//fe := e.new((tree.Ast)(fn.Body.Nodes))
	//_ = fe.run()
	// retV, retOk := fe.memory.GetAtTop("0")
	if evt != nil && evt.OnRelease != nil {
		evt.OnRelease()
	}

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
				retV[idx] = e.TypeCast(retT, retV[idx])
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
