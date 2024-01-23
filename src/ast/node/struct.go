package node

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type StructBlock struct {
	Token token.Token
	Body  *orderedmap.OrderedMap[string, *Declare] // map[identifier]type
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

func (n *StructBlock) GetToken() token.Token {
	return n.Token
}

type StructLiteral struct {
	Token token.Token
	Body  *orderedmap.OrderedMap[string, Expr]
	Type  *Type
}

func (n *StructLiteral) expr() {}

func (n *StructLiteral) TokenValue() string {
	return n.Token.Value
}

func (n *StructLiteral) GetToken() token.Token {
	return n.Token
}
