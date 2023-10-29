package node

import "git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"

type DeclareStmt struct {
	Type token.Token
	Name *Identifier
}

func (a *DeclareStmt) TokenValue() string {
	return "DeclareStmt"
}

func (a *DeclareStmt) stmt() {}

var _ Stmt = (*DeclareStmt)(nil)
