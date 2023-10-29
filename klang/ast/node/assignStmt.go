package node

import "git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"

type AssignStmt struct {
	Lhs struct {
		Type token.Token
		Name *Identifier
	}
	Rhs struct {
		Type token.Token
	}
}

func (a *AssignStmt) TokenValue() string {
	return "AssignStmt"
}

func (a *AssignStmt) stmt() {}

var _ Stmt = (*AssignStmt)(nil)
