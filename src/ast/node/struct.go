package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type StructBlock struct {
	Token token.Token
	Body  map[string]*Type // map[identifier]type
	Name  string
}

func (n *StructBlock) stmt() {}

func (n *StructBlock) TokenValue() string {
	return n.Token.Value
}

type StructLiteral struct {
	Token token.Token
	Body  map[string]Expr
}

func (n *StructLiteral) expr() {}

func (n *StructLiteral) TokenValue() string {
	return n.Token.Value
}
