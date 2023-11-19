package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
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
	switch scope.(type) {
	case *Eval:
		_e := scope.(*Eval)
		return _e.EvalFuncCall(funcCall)
		// TODO: More situation!
	}
	return nil
}

func (e *Eval) EvalExpr(n node.Expr) any {
	switch n.(type) {
	case *node.DotExpr:
		return e.EvalDotExpr(n.(*node.DotExpr))
	case *node.BinaryOperExpr:
		return e.EvalBinOperExpr(n.(*node.BinaryOperExpr))
	case *node.UnaryOperExpr:
		return e.EvalUnaryExpr(n.(*node.UnaryOperExpr))
	case *node.IntLiteral:
		return e.EvalIntLiteral(n.(*node.IntLiteral))
	case *node.FloatLiteral:
		return e.EvalFloatLiteral(n.(*node.FloatLiteral))
	case *node.StringLiteral:
		return e.EvalStringLiteral(n.(*node.StringLiteral))
	case *node.BoolLiteral:
		return e.EvalBoolLiteral(n.(*node.BoolLiteral))
	case *node.Identifier:
		return e.EvalIdentifier(n.(*node.Identifier))
	case *node.ArrayLiteral:
		return e.EvalArrayLiteral(n.(*node.ArrayLiteral))
	case *node.MapLiteral:
		return e.EvalMapLiteral(n.(*node.MapLiteral))
	case *node.FuncCall:
		return e.EvalFuncCall(n.(*node.FuncCall))
	case *node.LambdaExpr:
		return e.EvalLambdaExpr(n.(*node.LambdaExpr))
	case *node.AssignStmt:
		return e.EvalAssignStmt(n.(*node.AssignStmt))
	default:
		fmt.Println(reflect.TypeOf(n))
		panic("not implemented")
	}
}
