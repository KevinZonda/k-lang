package token

import "github.com/antlr4-go/antlr/v4"

type Kind int

type Token struct {
	Kind                   Kind
	Value                  string
	BeginLine, BeginColumn int
	EndLine, EndColumn     int
	Len                    int
}

func FromAntlrToken(t antlr.Token) Token {
	l := len(t.GetText())
	return Token{
		Kind:        Kind(t.GetTokenType()),
		Value:       t.GetText(),
		BeginLine:   t.GetLine(),
		BeginColumn: t.GetColumn(),
		Len:         l,
		EndLine:     t.GetLine(),
		EndColumn:   t.GetColumn() + l,
	}
}

func (t Token) WithBegin(begin antlr.Token) Token {
	t.BeginLine = begin.GetLine()
	t.BeginColumn = begin.GetColumn()
	t.EndLine = begin.GetLine()
	t.EndColumn = begin.GetColumn() + t.Len
	return t
}

func (t Token) WithEnd(end antlr.Token) Token {
	if end == nil {
		return t
	}
	endL := end.GetLine()
	endCol := end.GetColumn() + len(end.GetText())
	if endL > t.EndLine {
		t.EndLine = endL
		t.EndColumn = endCol
		return t
	}
	if endL == t.EndLine && endCol > t.EndColumn {
		t.EndColumn = endCol
	}
	return t
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
	Try
	Catch
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
