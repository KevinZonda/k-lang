package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
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

func exprVal(v any) ExprResult {
	return ExprResult{HasValue: true, Value: v}
}

func exprNoVal() ExprResult {
	return ExprResult{HasValue: false}
}

func (e *Eval) evalExpr(n node.Expr, keepRef bool) ExprResult {
	e.currentToken = n.GetToken()
	switch expr := n.(type) {
	case *node.DotExpr:
		return e.EvalDotExpr(expr)
	case *node.BinaryOperExpr:
		return exprVal(e.EvalBinOperExpr(expr))
	case *node.UnaryOperExpr:
		return exprVal(e.EvalUnaryExpr(expr))
	case *node.IntLiteral:
		return exprVal(e.EvalIntLiteral(expr))
	case *node.FloatLiteral:
		return exprVal(e.EvalFloatLiteral(expr))
	case *node.RefExpr:
		return e.EvalRefExpr(expr)
	case *node.StringLiteral:
		return exprVal(e.EvalStringLiteral(expr))
	case *node.BoolLiteral:
		return exprVal(e.EvalBoolLiteral(expr))
	case *node.Identifier:
		return exprVal(e.EvalIdentifier(expr, keepRef))
	case *node.ArrayLiteral:
		return exprVal(e.EvalArrayLiteral(expr))
	case *node.CommaExpr:
		return exprVal(e.EvalCommaExpr(expr))
	case *node.MapLiteral:
		return exprVal(e.EvalMapLiteral(expr))
	case *node.FuncCall:
		return e.EvalFuncCall(expr)
	case *node.LambdaExpr:
		return exprVal(e.EvalLambdaExpr(expr))
	case *node.AssignStmt:
		e.EvalAssignStmt(expr)
		return exprNoVal()
	case *node.StructLiteral:
		return exprVal(e.EvalStructLiteral(expr))
	case *node.IndexExpr:
		return exprVal(e.EvalIndexExpr(expr))
	case *node.NilLiteral:
		return exprVal(nil)
	default:
		panic("not implemented: " + reflect.TypeOf(n).String())
	}

}

func (e *Eval) EvalExpr(n node.Expr) ExprResult {
	return e.evalExpr(n, false)
}

func (e *Eval) EvalRefExpr(n *node.RefExpr) ExprResult {
	v := e.evalExpr(n.Expr, true)
	switch vT := v.Value.(type) {
	case *obj.Object:
		return exprVal(vT.CreateRef())
	default:
		return v
	}
}
