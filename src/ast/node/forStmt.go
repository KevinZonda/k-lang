package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type CStyleFor struct {
	baseStmt
	Token         token.Token
	InitialExpr   Expr
	ConditionExpr Expr
	AfterIterExpr Expr
	Body          *CodeBlock
}

func (c *CStyleFor) TokenValue() string {
	return c.Token.Value
}

type WhileStyleFor struct {
	baseStmt
	Token         token.Token
	ConditionExpr Expr
	Body          *CodeBlock
}

func (c *WhileStyleFor) TokenValue() string {
	return c.Token.Value
}

type IterStyleFor struct {
	baseStmt
	Token    token.Token
	Type     *Identifier
	Variable *Identifier
	Iterator Expr
	Body     *CodeBlock
}

func (c *IterStyleFor) TokenValue() string {
	return c.Token.Value
}
