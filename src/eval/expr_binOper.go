package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
)

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
