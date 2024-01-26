package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type ReturnStmt struct {
	Token token.Token
	Value []Expr
}

func (r *ReturnStmt) stmt() {}

func (r *ReturnStmt) TokenValue() string { return "return" }

func (r *ReturnStmt) GetToken() token.Token {
	return r.Token
}

type ContinueStmt struct {
	Token token.Token
}

func (c *ContinueStmt) stmt() {}

func (c *ContinueStmt) TokenValue() string { return "continue" }
func (c *ContinueStmt) GetToken() token.Token {
	return c.Token
}

type BreakStmt struct {
	Token token.Token
}

func (b *BreakStmt) stmt() {}

func (b *BreakStmt) TokenValue() string { return "break" }
func (b *BreakStmt) GetToken() token.Token {
	return b.Token
}
