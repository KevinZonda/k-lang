package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/memory"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
	"strings"
)

func (e *Eval) EvalDotExpr(n *node.DotExpr) ExprResult {
	var left any

	if lfc, ok := n.Left.(*node.FuncCall); ok {
		left = e.EvalFuncCall(lfc).EnsureValue()
	} else {
		left = e.EvalExpr(n.Left).EnsureValue()
	}

	e.currentToken = n.GetToken()

	if rightT, ok := n.Right.(*node.FuncCall); ok {
		return e.EvalFuncCallAfterScope(left, rightT)
	} else {
		return e.EvalPropertyAfterScope(left, n.Right)
	}
}

func (e *Eval) EvalPropertyAfterScope(scope any, property node.Expr) ExprResult {
	var actualPpt string
	switch properT := property.(type) {
	case *node.Identifier:
		actualPpt = properT.Value
	case *node.DotExpr:
		actualPpt = asType[string](e.EvalDotExpr(properT).EnsureValue())
	default:
		panic("Not Implemented EvalPropertyAfterScope" + reflect.TypeOf(property).String())
	}
	scope, _ = e.unboxObj(scope)
	switch leftT := scope.(type) {
	case *obj.StructField:
		if leftT.ParentEval != nil && leftT.ParentEval != e {
			if strings.HasPrefix(actualPpt, "_") {
				panic("No Access To Private Property: " + actualPpt)
			}
		}
		res, ok := leftT.Fields.Get(actualPpt)
		return ExprResult{HasValue: ok, Value: res}
	case *Eval:
		if leftT != e && strings.HasPrefix(actualPpt, "_") {
			panic("No Access To Private Function: " + actualPpt)
		}
		if o, ok := leftT.memory.Bottom().Get(actualPpt); ok {
			v, _ := leftT.unboxObj(o)
			return exprVal(v)
		}
		panic("No Property Found: " + actualPpt)
	default:
		panic("Not Implemented EvalPropertyAfterScope " + reflect.TypeOf(scope).String())
	}
	return exprNoVal()
}

func (e *Eval) EvalFuncCallAfterScope(scope any, funcCall *node.FuncCall) ExprResult {
	var _fc node.FuncCall
	_fc.Caller = funcCall.Caller
	_fc.Args = funcCall.Args
	_fc.Token = funcCall.Token
	scope, _ = e.unboxObj(scope)

	switch scopeT := scope.(type) {
	case *Eval:
		if scopeT != e && strings.HasPrefix(funcCall.Caller.Value, "_") {
			panic("No Access To Private Function: " + funcCall.Caller.Value)
		}

		return scopeT.EvalFuncCall(funcCall)
	case obj.ILibrary:
		v := scopeT.FuncCall(e.std, _fc.Caller.Value, e.evalExprs(_fc.Args...))
		return ExprResult{HasValue: v.HasValue(), Value: v.Value()}
	case *obj.StructField:
		if scopeT.ParentEval != nil && scopeT.ParentEval != e {
			if strings.HasPrefix(funcCall.Caller.Value, "_") {
				panic(fmt.Sprint("No Access To Private Function :"+funcCall.Caller.Value, ", Required: ", scopeT.ParentEval, " But Got: ", e))
			}
		}

		var funcB *node.FuncBlock

		v, ok := scopeT.Fields.Get(funcCall.Caller.Value)
		if !ok {
			panic("No Value Found From Struct: " + funcCall.Caller.Value)
		}
		switch vT := v.(type) {
		case *node.LambdaExpr:
			funcB = vT.ToFunc(funcCall.Caller.Value)
		case *node.FuncBlock:
			funcB = vT
		}
		if funcB == nil {
			panic("Func Nil Found From Struct: " + funcCall.Caller.Value)
		}
		if scopeT.ParentEval != nil {
			scopeE := scopeT.ParentEval.(*Eval)
			return scopeE.EvalFuncBlock(funcB, _fc.Args,
				&EvalFuncBlockEvent{
					OnNewFrameCreated: func(topF *memory.Layer) {
						topF.SetValue("self", scopeT)
					},
				})
		}

		return e.EvalFuncBlock(funcB, _fc.Args, &EvalFuncBlockEvent{
			OnNewFrameCreated: func(topF *memory.Layer) {
				topF.SetValue("self", scopeT)
			},
		})
	// TODO: More situation!
	case string:
		v := e.builtInString(scopeT, _fc)
		return ExprResult{HasValue: v != nil, Value: v}
	case []any:
		v := e.builtInArray(scopeT, _fc)
		return ExprResult{HasValue: v != nil, Value: v}
	case map[any]any:
		v := e.builtInMap(scopeT, _fc)
		return ExprResult{HasValue: v != nil, Value: v}
	default:
		panic("EvalFuncCallAfterScope Not Implemented" + reflect.TypeOf(scope).String())
	}
	return exprNoVal()
}
