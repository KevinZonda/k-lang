package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type StructBlock struct {
	Token token.Token
	Body  map[string]*Declare // map[identifier]type
	Name  string
}

type Declare struct {
	Type  *Type
	Value Expr
}

func (n *StructBlock) block() {}

func (n *StructBlock) TokenValue() string {
	return n.Token.Value
}

type StructLiteral struct {
	Token token.Token
	Body  map[string]Expr
	Type  *Type
}

func (n *StructLiteral) expr() {}

func (n *StructLiteral) TokenValue() string {
	return n.Token.Value
}
