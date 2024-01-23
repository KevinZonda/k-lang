package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type IfStmt struct {
	Token     token.Token
	Condition Expr
	IfTrue    *CodeBlock
	IfFalse   *CodeBlock
}

func (i *IfStmt) TokenValue() string {
	return i.Token.Value
}

func (i *IfStmt) stmt() {}

func (i *IfStmt) GetToken() token.Token {
	return i.Token
}
