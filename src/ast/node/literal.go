package node

import (
	"fmt"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
)

type ILiteralValue interface {
	ConstVal() any
}

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

func (i *IntLiteral) ConstVal() any {
	return i.Value
}

func (i *IntLiteral) GetToken() token.Token {
	return i.Token
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

func (f *FloatLiteral) ConstVal() any {
	return f.Value
}

func (f *FloatLiteral) GetToken() token.Token {
	return f.Token
}

//endregion

//region String

type StringLiteral struct {
	Token token.Token
	Value string
	Mode  uint8
	Char  uint8 // " or ' as begin and end?
}

func (s *StringLiteral) expr() {}

func (s *StringLiteral) TokenValue() string {
	return s.Value
}

func (s *StringLiteral) String() string {
	return fmt.Sprintf("(%s:string)", s.Value)
}

func (s *StringLiteral) ConstVal() any {
	return s.Value
}

func (s *StringLiteral) GetToken() token.Token {
	return s.Token
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
	return fmt.Sprintf("(%tx:bool)", b.Value)
}

func (b *BoolLiteral) ConstVal() any {
	return b.Value
}

func (b *BoolLiteral) GetToken() token.Token {
	return b.Token
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

func (a *ArrayLiteral) GetToken() token.Token {
	return a.Token
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

func (m *MapPairLiteral) GetToken() token.Token {
	return m.Token
}

type MapLiteral struct {
	Token token.Token
	Value []*MapPairLiteral
}

func (m *MapLiteral) expr() {}

func (m *MapLiteral) TokenValue() string {
	return m.Token.Value
}

func (m *MapLiteral) GetToken() token.Token {
	return m.Token
}
