package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
)

func (e *Eval) EvalExpr(n node.Expr) any {
	e.currentToken = n.GetToken()
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
		e.EvalAssignStmt(n.(*node.AssignStmt))
		return nil
	case *node.StructLiteral:
		return e.EvalStructLiteral(n.(*node.StructLiteral))
	case *node.IndexExpr:
		return e.EvalIndexExpr(n.(*node.IndexExpr))
	default:
		fmt.Println(reflect.TypeOf(n))
		panic("not implemented")
	}
}
