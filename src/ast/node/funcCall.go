package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type FuncCall struct {
	Token  token.Token
	Caller *Identifier

	CallExpr Expr // TODO:

	Args []Expr
}

func (f *FuncCall) TokenValue() string {
	return f.Token.Value
}
func (f *FuncCall) stmt() {}
func (f *FuncCall) expr() {}
func (f *FuncCall) GetToken() token.Token {
	return f.Token
}
