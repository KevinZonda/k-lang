package node

import (
	"fmt"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
)

//region Int

type IntLiteral struct {
	Token token.Token
	Value int
}

func (i *IntLiteral) expr() {}

func (i *IntLiteral) TokenValue() string {
	return fmt.Sprint(i.Value)
}

func (i *IntLiteral) String() string {
	return fmt.Sprintf("(%d:int)", i.Value)
}

//endregion

//region Float

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (f *FloatLiteral) expr() {}

func (f *FloatLiteral) TokenValue() string {
	return f.Token.Value
}

func (f *FloatLiteral) String() string {
	return fmt.Sprintf("(%f:float)", f.Value)
}

//endregion

//region String

type StringLiteral struct {
	Token token.Token
	Value string
}

func (s *StringLiteral) expr() {}

func (s *StringLiteral) TokenValue() string {
	return s.Value
}

func (s *StringLiteral) String() string {
	return fmt.Sprintf("(%s:string)", s.Value)
}

//endregion

//region Bool

type BoolLiteral struct {
	Token token.Token
	Value bool
}

func (b *BoolLiteral) expr() {}

func (b *BoolLiteral) TokenValue() string {
	return fmt.Sprint(b.Value)
}

func (b *BoolLiteral) String() string {
	return fmt.Sprintf("(%t:bool)", b.Value)
}

//endregion

//region ArrayLiteral

type ArrayLiteral struct {
	Token token.Token
	Value []Expr
}

func (a *ArrayLiteral) expr() {}

func (a *ArrayLiteral) TokenValue() string {
	return a.Token.Value
}

//endregion

type MapPairLiteral struct {
	Key   Expr
	Value Expr
	Token token.Token
}

func (m *MapPairLiteral) expr() {}

func (m *MapPairLiteral) TokenValue() string {
	return m.Token.Value
}

type MapLiteral struct {
	Token token.Token
	Value []*MapPairLiteral
}

func (m *MapLiteral) expr() {}

func (m *MapLiteral) TokenValue() string {
	return m.Token.Value
}
