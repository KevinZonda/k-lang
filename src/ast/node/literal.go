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

func (i *IntLiteral) TokenValue() string {
	return fmt.Sprint(i.Value)
}

func (i *IntLiteral) expr() {}

func (i *IntLiteral) String() string {
	return fmt.Sprintf("(%d:int)", i.Value)
}

//endregion

var _ Expr = (*IntLiteral)(nil)

//region Float

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (f *FloatLiteral) TokenValue() string {
	return f.Token.Value
}

func (f *FloatLiteral) expr() {}

var _ Expr = (*FloatLiteral)(nil)

func (f *FloatLiteral) String() string {
	return fmt.Sprintf("(%f:float)", f.Value)
}

//endregion

//region String

type StringLiteral struct {
	Token token.Token
	Value string
}

func (s *StringLiteral) TokenValue() string {
	return s.Value
}

func (s *StringLiteral) expr() {}

var _ Expr = (*StringLiteral)(nil)

func (s *StringLiteral) String() string {
	return fmt.Sprintf("(%s:string)", s.Value)
}

//endregion

//region Bool

type BoolLiteral struct {
	Token token.Token
	Value bool
}

func (b *BoolLiteral) TokenValue() string {
	return fmt.Sprint(b.Value)
}

func (b *BoolLiteral) expr() {}

func (b *BoolLiteral) String() string {
	return fmt.Sprintf("(%t:bool)", b.Value)
}

var _ Expr = (*BoolLiteral)(nil)

//endregion

type ArrayLiteral struct {
	Token token.Token
	Value []Expr
}

func (a *ArrayLiteral) TokenValue() string {
	return a.Token.Value
}

func (a *ArrayLiteral) expr() {}

var _ Expr = (*ArrayLiteral)(nil)
