package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) EvalBinOperExpr(n *node.BinaryOperExpr) any {
	left := e.EvalExpr(n.Left).EnsureValue()
	if n.Token.Kind == token.Or && left == true {
		return true
	}
	if n.Token.Kind == token.And && left == false {
		return false
	}
	right := e.EvalExpr(n.Right).EnsureValue()
	left = obj.UnboxToEnd(left)
	right = obj.UnboxToEnd(right)
	e.currentToken = n.GetToken()
	return binaryOperEval.BinaryOper(n.Token.Kind, left, right)
}
