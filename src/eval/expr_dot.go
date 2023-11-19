package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/builtin"
)

func (e *Eval) EvalDotExpr(n *node.DotExpr) any {
	var left any

	if lfc, ok := n.Left.(*node.FuncCall); ok {
		left = e.EvalFuncCall(lfc)
	} else {
		left = e.EvalExpr(n.Left)
	}

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
	}
	fmt.Println("Scope: ", actualPpt)
	fmt.Println("Property: ", property)
	return nil

}

func (e *Eval) EvalFuncCallAfterScope(scope any, funcCall *node.FuncCall) any {
	var _fc node.FuncCall
	_fc.Caller = funcCall.Caller
	_fc.Args = funcCall.Args
	_fc.Token = funcCall.Token
	switch scope.(type) {
	case *Eval:
		_e := scope.(*Eval)
		return _e.EvalFuncCall(funcCall)
	case builtin.ILibrary:
		_lib := scope.(builtin.ILibrary)
		var args []any
		for _, expr := range _fc.Args {
			args = append(args, e.EvalExpr(expr))
		}
		return _lib.FuncCall(_fc.Caller.Value, args)
		// TODO: More situation!
	}
	return nil
}
