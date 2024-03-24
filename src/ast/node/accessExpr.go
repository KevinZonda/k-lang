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
	Token    token.Token
	Left     Expr
	Index    Expr
	EndIndex Expr
	Col      bool
}

func (a *IndexExpr) expr() {}
func (a *IndexExpr) TokenValue() string {
	return a.Token.Value
}
func (a *IndexExpr) GetToken() token.Token {
	return a.Token
}

type RefExpr struct {
	Token token.Token
	Expr  Expr
}

func (a *RefExpr) expr() {}
func (a *RefExpr) TokenValue() string {
	return a.Token.Value
}
func (a *RefExpr) GetToken() token.Token {
	return a.Token
}
