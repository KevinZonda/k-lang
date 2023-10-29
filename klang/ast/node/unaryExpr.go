package node

import (
	"fmt"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"
)

type UnaryExpr struct {
	Token token.Token
	Oper  string
	Expr  Expr
}

func (u *UnaryExpr) TokenValue() string {
	return u.Oper
}

func (u *UnaryExpr) expr() {}

func (u *UnaryExpr) String() string {
	return "(" + u.Oper + fmt.Sprint(u.Expr) + ")"
}

var _ Expr = (*UnaryExpr)(nil)
