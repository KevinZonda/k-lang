package token

import "github.com/antlr4-go/antlr/v4"

type Kind int

type Token struct {
	Kind  Kind
	Value string
}

func FromAntlrToken(t antlr.Token) Token {
	return Token{
		Kind:  Kind(t.GetTokenType()),
		Value: t.GetText(),
	}
}

const (
	LBrack Kind = iota + 1
	RBrack
	LParen
	RParen
	LSquare
	RSquare
	Comma
	Col
	Semi
	BinaryOper
	Equals
	NotEq
	GreaterEq
	LessEq
	Greater
	Less
	Or
	And
	Add
	Sub
	Mul
	Div
	Mod
	Dot
	To
	Struct
	Map
	Function
	Return
	Case
	Default
	Open
	As
	If
	Else
	For
	Match
	Break
	Continue
	IntegerLiteral
	NumberLiteral
	StringLiteral
	Not
	Assign
	Identity
	WS
	NewLine
)
