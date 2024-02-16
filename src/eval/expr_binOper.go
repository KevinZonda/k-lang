package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
)

func (e *Eval) EvalBinOperExpr(n *node.BinaryOperExpr) any {
	left := e.EvalExpr(n.Left).EnsureValue()
	right := e.EvalExpr(n.Right).EnsureValue()
	e.currentToken = n.GetToken()
	return binaryOperEval.BinaryOper(n.Token.Kind, left, right)
}
