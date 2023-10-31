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
	EOF         = antlr.TokenEOF
	LBrack Kind = iota
	RBrack
	LParen
	RParen
	LSquare
	RSquare
	Comma
	Col
	Semi
	Equals
	NotEq
	GreaterEq
	LessEq
	Greater
	Less
	Or
	And
	Pow
	Mul
	Div
	Add
	Sub
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
	True
	False
	IntegerLiteral
	NumberLiteral
	StringLiteral
	Not
	Assign
	Identifier
	Comment
	WS
	NewLine
)
