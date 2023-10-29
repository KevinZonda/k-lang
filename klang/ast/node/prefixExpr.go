package node

import "git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"

type PrefixExpr struct {
	Token token.Token
	Oper  string
	Right Expr
}

func (p *PrefixExpr) TokenValue() string {
	return p.Oper
}

func (p *PrefixExpr) expr() {}

var _ Expr = (*PrefixExpr)(nil)
