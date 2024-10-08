package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type AssignStmt struct {
	Token token.Token

	Assignee []*Assignee

	Value Expr
}

func (a *AssignStmt) TokenValue() string {
	return "AssignStmt"
}

func (a *AssignStmt) stmt() {}
func (a *AssignStmt) GetToken() token.Token {
	return a.Token
}

type Assignee struct {
	Type *Type
	Var  *Variable
}
