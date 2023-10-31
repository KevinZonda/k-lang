package node

import (
	"fmt"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/ast/token"
)

type UnaryOperExpr struct {
	Token token.Token
	Oper  string
	Expr  Expr
}

func (u *UnaryOperExpr) TokenValue() string {
	return u.Oper
}

func (u *UnaryOperExpr) expr() {}

func (u *UnaryOperExpr) String() string {
	return "(" + u.Oper + fmt.Sprint(u.Expr) + ")"
}

var _ Expr = (*UnaryOperExpr)(nil)
