package node

import (
	"fmt"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
)

type BinaryOperExpr struct {
	Token token.Token
	Left  Expr
	Oper  string
	Right Expr
}

func (a *BinaryOperExpr) expr() {}

func (i *BinaryOperExpr) TokenValue() string {
	return i.Oper
}

func (i *BinaryOperExpr) String() string {
	return fmt.Sprintf("{LHS=%s\nOper=%s\nRHS=%s}\n", i.Left, i.Oper, i.Right)
}
func (i *BinaryOperExpr) GetToken() token.Token {
	return i.Token
}

var _ Expr = (*BinaryOperExpr)(nil)
