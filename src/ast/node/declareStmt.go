package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type DeclareStmt struct {
	baseStmt
	Type token.Token
	Name *Identifier
}

func (a *DeclareStmt) TokenValue() string {
	return "DeclareStmt"
}
