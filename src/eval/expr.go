package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
)

func (e *Eval) evalExpr(n node.Expr, keepRef bool) any {
	e.currentToken = n.GetToken()
	switch expr := n.(type) {
	case *node.DotExpr:
		return e.EvalDotExpr(expr)
	case *node.BinaryOperExpr:
		return e.EvalBinOperExpr(expr)
	case *node.UnaryOperExpr:
		return e.EvalUnaryExpr(expr)
	case *node.IntLiteral:
		return e.EvalIntLiteral(expr)
	case *node.FloatLiteral:
		return e.EvalFloatLiteral(expr)
	case *node.StringLiteral:
		return e.EvalStringLiteral(expr)
	case *node.BoolLiteral:
		return e.EvalBoolLiteral(expr)
	case *node.Identifier:
		return e.EvalIdentifier(expr, keepRef)
	case *node.ArrayLiteral:
		return e.EvalArrayLiteral(expr)
	case *node.CommaExpr:
		return e.EvalCommaExpr(expr)
	case *node.MapLiteral:
		return e.EvalMapLiteral(expr)
	case *node.FuncCall:
		return e.EvalFuncCall(expr)
	case *node.LambdaExpr:
		return e.EvalLambdaExpr(expr)
	case *node.AssignStmt:
		e.EvalAssignStmt(expr)
		return nil
	case *node.StructLiteral:
		return e.EvalStructLiteral(expr)
	case *node.IndexExpr:
		return e.EvalIndexExpr(expr)
	case *node.NilLiteral:
		return nil
	default:
		panic("not implemented: " + reflect.TypeOf(n).String())
	}

}

func (e *Eval) EvalExpr(n node.Expr) any {
	return e.evalExpr(n, false)
}
