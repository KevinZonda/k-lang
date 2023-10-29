package node

import (
	"fmt"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"
)

type InfixExpr struct {
	Token token.Token
	Left  Expr
	Oper  string
	Right Expr
}

func (i *InfixExpr) TokenValue() string {
	return i.Oper
}

func (i *InfixExpr) expr() {}

func (i *InfixExpr) String() string {
	return fmt.Sprintf("{LHS=%s\nOper=%s\nRHS=%s}\n", i.Left, i.Oper, i.Right)
}

var _ Expr = (*InfixExpr)(nil)
