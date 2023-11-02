package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type FuncCall struct {
	baseExpr
	baseStmt
	Token  token.Token
	Caller *Variable
	Args   []Expr
}

func (f *FuncCall) TokenValue() string {
	return f.Token.Value
}
