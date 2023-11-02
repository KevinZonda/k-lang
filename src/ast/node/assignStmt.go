package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type AssignStmt struct {
	baseStmt
	Token token.Token

	Type *Identifier
	Var  *Variable

	Value Expr
}

func (a *AssignStmt) TokenValue() string {
	return "AssignStmt"
}
