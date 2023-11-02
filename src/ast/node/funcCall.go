package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type FuncCall struct {
	Token  token.Token
	Caller *Variable
	Args   []Expr
}

func (f *FuncCall) expr() {}
func (f *FuncCall) stmt() {}

func (f *FuncCall) TokenValue() string {
	return f.Token.Value
}

var _ Expr = (*FuncCall)(nil)
var _ Stmt = (*FuncCall)(nil)
