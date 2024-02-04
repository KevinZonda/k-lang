package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
	"strings"
)

func (e *Eval) EvalDotExpr(n *node.DotExpr) any {
	var left any

	if lfc, ok := n.Left.(*node.FuncCall); ok {
		left = e.EvalFuncCall(lfc)
	} else {
		left = e.EvalExpr(n.Left)
	}

	e.currentToken = n.GetToken()

	if rightT, ok := n.Right.(*node.FuncCall); ok {
		return e.EvalFuncCallAfterScope(left, rightT)
	} else {
		return e.EvalPropertyAfterScope(left, n.Right)
	}
}

func (e *Eval) EvalPropertyAfterScope(scope any, property node.Expr) any {
	var actualPpt string
	switch properT := property.(type) {
	case *node.Identifier:
		actualPpt = properT.Value
	case *node.DotExpr:
		actualPpt = e.EvalDotExpr(properT).(string)
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
		if !ok {
			res = nil
		}
		return res
	case *Eval:
		if leftT != e && strings.HasPrefix(actualPpt, "_") {
			panic("No Access To Private Function: " + actualPpt)
		}
		if o, ok := leftT.objTable.Bottom().Get(actualPpt); ok {
			v, _ := leftT.unboxObj(o)
			return v
		}
		panic("No Property Found: " + actualPpt)
	default:
		panic("Not Implemented EvalPropertyAfterScope " + reflect.TypeOf(scope).String())
	}
	return nil

}

func (e *Eval) EvalFuncCallAfterScope(scope any, funcCall *node.FuncCall) any {
	var _fc node.FuncCall
	_fc.Caller = funcCall.Caller
	_fc.Args = funcCall.Args
	_fc.Token = funcCall.Token

	switch scopeT := scope.(type) {
	case *Eval:
		if scopeT != e && strings.HasPrefix(funcCall.Caller.Value, "_") {
			panic("No Access To Private Function: " + funcCall.Caller.Value)
		}

		return scopeT.EvalFuncCall(funcCall)
	case obj.ILibrary:
		return scopeT.FuncCall(e.builtin, _fc.Caller.Value, e.evalExprs(_fc.Args...))
	case *obj.StructField:
		if scopeT.ParentEval != nil && scopeT.ParentEval != e {
			if strings.HasPrefix(funcCall.Caller.Value, "_") {
				panic("No Access To Private Function: " + funcCall.Caller.Value)
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

		return e.EvalFuncBlock(funcB, _fc.Args, func() {
			e.objTable.SetAtTop("self", scopeT)
		})
		// TODO: More situation!
	default:
		panic("EvalFuncCallAfterScope Not Implemented" + reflect.TypeOf(scope).String())
	}
	return nil
}
