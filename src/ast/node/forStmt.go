package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type CStyleFor struct {
	baseStmt
	Token         token.Token
	InitialExpr   Expr
	ConditionExpr Expr
	AfterIterExpr Expr
}

func (c *CStyleFor) TokenValue() string {
	return c.Token.Value
}
