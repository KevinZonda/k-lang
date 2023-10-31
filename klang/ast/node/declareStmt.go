package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/ast/token"

type DeclareStmt struct {
	Type token.Token
	Name *Identifier
}

func (a *DeclareStmt) TokenValue() string {
	return "DeclareStmt"
}

func (a *DeclareStmt) stmt() {}

var _ Stmt = (*DeclareStmt)(nil)
