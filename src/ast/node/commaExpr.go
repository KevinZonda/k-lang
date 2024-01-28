package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type CommaExpr struct {
	Token token.Token
	Exprs []Expr
}

func (a *CommaExpr) expr() {}
func (a *CommaExpr) TokenValue() string {
	return a.Token.Value
}
func (a *CommaExpr) GetToken() token.Token {
	return a.Token
}
