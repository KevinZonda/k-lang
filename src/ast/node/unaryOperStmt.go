package node

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
)

type UnaryOperStmt struct {
	Token      token.Token
	Oper       string
	Identifier *Identifier
}

func (u *UnaryOperStmt) TokenValue() string {
	return u.Oper
}

func (u *UnaryOperStmt) String() string {
	return "(" + u.Oper + fmt.Sprint(u.Identifier) + ")"
}

func (u *UnaryOperStmt) stmt() {}

func (u *UnaryOperStmt) GetToken() token.Token {
	return u.Token
}
