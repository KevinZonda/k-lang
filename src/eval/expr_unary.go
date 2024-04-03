package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

func (e *Eval) EvalUnaryExpr(n *node.UnaryOperExpr) any {
	val := e.EvalExpr(n.Expr).EnsureValue()
	val = obj.UnboxToEnd(val)
	e.currentToken = n.GetToken()

	factor := 1
	//if n.Token.Kind == parser.V2LexerSub || n.Token.Kind == parser.V2LexerNot {
	if n.Token.Kind == token.Sub || n.Token.Kind == token.Not {
		factor = -1
	}

	switch valT := val.(type) {
	case float64:
		return valT * float64(factor)
	case int:
		return valT * factor
	case bool:
		if factor == 1 {
			return val
		}
		return !valT
	default:
		panic("eval unary expr not supported type: " + reflect.TypeOf(val).String())
	}
}

func (e *Eval) EvalTrinaryOperExpr(n *node.TrinaryOperExpr) any {
	e.currentToken = n.GetToken()
	cond := e.EvalExpr(n.Cond).EnsureValue().(bool)
	if cond {
		return e.EvalExpr(n.IfTrue).EnsureValue()
	}
	return e.EvalExpr(n.IfFalse).EnsureValue()
}
