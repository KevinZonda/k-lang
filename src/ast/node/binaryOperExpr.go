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

func (i *BinaryOperExpr) TokenValue() string {
	return i.Oper
}

func (i *BinaryOperExpr) expr() {}

func (i *BinaryOperExpr) String() string {
	return fmt.Sprintf("{LHS=%s\nOper=%s\nRHS=%s}\n", i.Left, i.Oper, i.Right)
}

var _ Expr = (*BinaryOperExpr)(nil)
