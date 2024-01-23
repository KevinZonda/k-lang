package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type DotExpr struct {
	Token token.Token
	Left  Expr
	Right Expr
}

func (a *DotExpr) expr() {}
func (a *DotExpr) TokenValue() string {
	return a.Token.Value
}
func (a *DotExpr) GetToken() token.Token {
	return a.Token
}

type IndexExpr struct {
	Token token.Token
	Left  Expr
	Index Expr
}

func (a *IndexExpr) expr() {}
func (a *IndexExpr) TokenValue() string {
	return a.Token.Value
}
func (a *IndexExpr) GetToken() token.Token {
	return a.Token
}
