package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type CStyleFor struct {
	Token         token.Token
	InitialStmt   Stmt
	ConditionExpr Expr
	AfterIterStmt Stmt
	Body          *CodeBlock
}

func (c *CStyleFor) TokenValue() string {
	return c.Token.Value
}

func (c *CStyleFor) GetToken() token.Token {
	return c.Token
}

func (c *CStyleFor) stmt() {}

type WhileStyleFor struct {
	Token         token.Token
	ConditionExpr Expr
	Body          *CodeBlock
}

func (c *WhileStyleFor) TokenValue() string {
	return c.Token.Value
}

func (c *WhileStyleFor) GetToken() token.Token {
	return c.Token
}

func (a *WhileStyleFor) stmt() {}

type IterStyleFor struct {
	Token    token.Token
	Type     *Type
	Variable *Identifier
	Iterator Expr
	Body     *CodeBlock
}

func (c *IterStyleFor) TokenValue() string {
	return c.Token.Value
}

func (c *IterStyleFor) GetToken() token.Token {
	return c.Token
}
func (a *IterStyleFor) stmt() {}
