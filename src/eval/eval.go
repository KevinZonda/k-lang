package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
)

type Eval struct {
	ast tree.Ast
	it  any
}

func (e *Eval) It() any {
	return e.it
}

func (e *Eval) Do() {
	for _, n := range e.ast {
		switch n.(type) {
		case node.Expr:
			e.it = e.EvalExpr(n.(node.Expr))
		default:
			panic("not implemented")
		}
	}
}

func (e *Eval) EvalBinOperExpr(n *node.BinaryOperExpr) any {
	left := e.EvalExpr(n.Left)
	right := e.EvalExpr(n.Right)

	switch n.Token.Kind {
	case token.Add:
		return binaryOperEval.Add(left, right)
	case token.Sub:
		return binaryOperEval.Sub(left, right)
	case token.Mul:
		return binaryOperEval.Mul(left, right)
	case token.Pow:
		return binaryOperEval.Pow(left, right)
	case token.Div:
		return binaryOperEval.Div(left, right)
	case token.Mod:
		return binaryOperEval.Mod(left, right)
	case token.And:
		return binaryOperEval.And(left, right)
	case token.Equals:
		return binaryOperEval.Eq(left, right)
	case token.NotEq:
		return !binaryOperEval.Eq(left, right)
	case token.Or:
		return binaryOperEval.Or(left, right)
	case token.Less:
		return binaryOperEval.Less(left, right)
	case token.LessEq:
		return binaryOperEval.LessEq(left, right)
	case token.Greater:
		return binaryOperEval.Greater(left, right)
	case token.GreaterEq:
		return binaryOperEval.GreaterEq(left, right)
	}
	panic("not implemented")
}

func (e *Eval) EvalUnaryExpr(n *node.UnaryOperExpr) any {
	val := e.EvalExpr(n.Expr)
	factor := 1
	if n.Token.Kind == token.Sub || n.Token.Kind == token.Not {
		factor = -1
	}

	switch val.(type) {
	case float64:
		return val.(float64) * float64(factor)
	case int:
		return val.(int) * factor
	case bool:
		if factor == 1 {
			return val
		}
		return !val.(bool)
	default:
		panic("not supported type")
	}
}

//	func ifNumber[T any](n any, ifFloat func(float64) T, ifInt func(int) T) (T, bool) {
//		switch n.(type) {
//		case float64:
//
//			return ifFloat(n.(float64)), true
//		case int:
//
//			return ifInt(n.(int)), true
//		default:
//			var r T
//			return r, false
//		}
//	}

func (e *Eval) EvalStringLiteral(n *node.StringLiteral) string {
	return n.Value
}

func (e *Eval) EvalFloatLiteral(n *node.FloatLiteral) float64 {
	return n.Value
}

func (e *Eval) EvalIntLiteral(n *node.IntLiteral) int {
	return n.Value
}

func (e *Eval) EvalBoolLiteral(n *node.BoolLiteral) bool {
	return n.Value
}

func (e *Eval) EvalExpr(n node.Expr) any {
	switch n.(type) {
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
	default:
		panic("not implemented")
	}
}

func New(ast tree.Ast) *Eval {
	return &Eval{
		ast: ast,
	}
}
