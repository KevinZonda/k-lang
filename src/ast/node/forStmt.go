package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type CStyleFor struct {
	Token         token.Token
	InitialExpr   Expr
	ConditionExpr Expr
	AfterIterExpr Expr
	Body          *CodeBlock
}

func (c *CStyleFor) TokenValue() string {
	return c.Token.Value
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

func (a *WhileStyleFor) stmt() {}

type IterStyleFor struct {
	Token    token.Token
	Type     *Identifier
	Variable *Identifier
	Iterator Expr
	Body     *CodeBlock
}

func (c *IterStyleFor) TokenValue() string {
	return c.Token.Value
}
func (a *IterStyleFor) stmt() {}
