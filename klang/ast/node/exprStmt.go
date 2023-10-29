package node

import "git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"

type ExprStmt struct {
	token.Token
	Expr Expr
}

func (e *ExprStmt) TokenValue() string {
	return "exprStmt"
}

func (e *ExprStmt) stmt() {}

var _ Stmt = (*ExprStmt)(nil)
