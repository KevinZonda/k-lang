package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type OpenStmt struct {
	Token token.Token
	Path  string
	As    string
}

func (n *OpenStmt) TokenValue() string {
	return n.Token.Value
}

func (n *OpenStmt) stmt() {}

func (n *OpenStmt) GetToken() token.Token {
	return n.Token
}

type OpenBlock struct {
	Token   token.Token
	Openers []*OpenStmt
}

func (o *OpenBlock) block() {}
func (o *OpenBlock) TokenValue() string {
	return o.Token.Value
}
func (o *OpenBlock) GetToken() token.Token {
	return o.Token
}
