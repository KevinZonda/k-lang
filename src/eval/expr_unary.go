package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
)

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
