package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type MatchStmt struct {
	Token   token.Token
	Match   Expr
	Cases   []*MatchCase
	Default *CodeBlock
}

func (m *MatchStmt) stmt() {}
func (m *MatchStmt) TokenValue() string {
	return m.Token.Value
}

func (m *MatchStmt) GetToken() token.Token {
	return m.Token
}

type MatchCase struct {
	Expr []Expr
	Body *CodeBlock
}
