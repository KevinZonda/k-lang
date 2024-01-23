package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type DeclareStmt struct {
	Token token.Token
	Name  *Identifier
	Type  *Type
	Value Expr
}

func (a *DeclareStmt) TokenValue() string {
	return "DeclareStmt"
}

func (a *DeclareStmt) GetToken() token.Token {
	return a.Token
}

func (a *DeclareStmt) stmt() {}
