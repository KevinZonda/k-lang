package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
)

func (e *Eval) EvalDotExpr(n *node.DotExpr) any {
	var left any

	if lfc, ok := n.Left.(*node.FuncCall); ok {
		left = e.EvalFuncCall(lfc)
	} else {
		left = e.EvalExpr(n.Left)
	}

	e.currentToken = n.GetToken()

	if _, ok := n.Right.(*node.FuncCall); ok {
		return e.EvalFuncCallAfterScope(left, n.Right.(*node.FuncCall))
	} else {
		return e.EvalPropertyAfterScope(left, n.Right)
	}
}

func (e *Eval) EvalPropertyAfterScope(scope any, property node.Expr) any {
	var actualPpt any
	switch property.(type) {
	case *node.Identifier:
		actualPpt = property.(*node.Identifier).Value
	case *node.DotExpr:
		actualPpt = e.EvalDotExpr(property.(*node.DotExpr))
	default:
		panic("Not Implemented EvalPropertyAfterScope" + reflect.TypeOf(property).String())
	}
	//fmt.Println("Scope: ", actualPpt)
	//fmt.Println("Property: ", property)
	scope, _ = e.unboxObj(scope)
	switch scope.(type) {
	case *obj.StructField:
		_sf := scope.(*obj.StructField)
		res, ok := _sf.Fields.Get(actualPpt.(string))
		if !ok {
			res = nil
		}
		return res
	case *Eval:
		_e := scope.(*Eval)
		if o, ok := _e.objTable.Bottom().Get(actualPpt.(string)); ok {
			v, _ := _e.unboxObj(o)
			return v
		}
		panic("No Property Found: " + actualPpt.(string))
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
		return scopeT.EvalFuncCall(funcCall)
	case obj.ILibrary:
		return scopeT.FuncCall(e.builtin, _fc.Caller.Value, e.evalExprs(_fc.Args...))
	case *obj.StructField:
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
