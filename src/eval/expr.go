package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
)

type ExprResult struct {
	HasValue bool
	Value    any
}

func (r ExprResult) EnsureValue() any {
	if !r.HasValue {
		panic("no value")
	}
	return r.Value
}

func (e *Eval) evalExpr(n node.Expr, keepRef bool) ExprResult {
	e.currentToken = n.GetToken()
	switch expr := n.(type) {
	case *node.DotExpr:
		return e.EvalDotExpr(expr)
	case *node.BinaryOperExpr:
		return ExprResult{HasValue: true, Value: e.EvalBinOperExpr(expr)}
	case *node.UnaryOperExpr:
		return ExprResult{HasValue: true, Value: e.EvalUnaryExpr(expr)}
	case *node.IntLiteral:
		return ExprResult{HasValue: true, Value: e.EvalIntLiteral(expr)}
	case *node.FloatLiteral:
		return ExprResult{HasValue: true, Value: e.EvalFloatLiteral(expr)}
	case *node.StringLiteral:
		return ExprResult{HasValue: true, Value: e.EvalStringLiteral(expr)}
	case *node.BoolLiteral:
		return ExprResult{HasValue: true, Value: e.EvalBoolLiteral(expr)}
	case *node.Identifier:
		return ExprResult{HasValue: true, Value: e.EvalIdentifier(expr, keepRef)}
	case *node.ArrayLiteral:
		return ExprResult{HasValue: true, Value: e.EvalArrayLiteral(expr)}
	case *node.CommaExpr:
		return ExprResult{HasValue: true, Value: e.EvalCommaExpr(expr)}
	case *node.MapLiteral:
		return ExprResult{HasValue: true, Value: e.EvalMapLiteral(expr)}
	case *node.FuncCall:
		return e.EvalFuncCall(expr)
	case *node.LambdaExpr:
		return ExprResult{HasValue: true, Value: e.EvalLambdaExpr(expr)}
	case *node.AssignStmt:
		e.EvalAssignStmt(expr)
		return ExprResult{HasValue: false}
	case *node.StructLiteral:
		return ExprResult{HasValue: true, Value: e.EvalStructLiteral(expr)}
	case *node.IndexExpr:
		return ExprResult{HasValue: true, Value: e.EvalIndexExpr(expr)}
	case *node.NilLiteral:
		return ExprResult{HasValue: true, Value: nil}
	default:
		panic("not implemented: " + reflect.TypeOf(n).String())
	}

}

func (e *Eval) EvalExpr(n node.Expr) ExprResult {
	return e.evalExpr(n, false)
}
