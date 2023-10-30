package node

import "git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"

type AssignStmt struct {
	Token token.Token
	Lhs   struct {
		Type *Identifier
		Name *Identifier
	}
	Rhs struct {
		Value Expr
	}
}

func (a *AssignStmt) TokenValue() string {
	return "AssignStmt"
}

func (a *AssignStmt) stmt() {}

var _ Stmt = (*AssignStmt)(nil)
