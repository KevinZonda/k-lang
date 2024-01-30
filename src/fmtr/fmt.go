package fmtr

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

type str struct {
	strings.Builder
}

func (s *str) EndsWith(str string) bool {
	return strings.HasSuffix(s.String(), str)
}

func (s *str) EndsWithSpaces() bool {
	ss := strings.Split(s.String(), "\n")
	return strings.TrimSpace(ss[len(ss)-1]) == ""
}

func (s *str) CountSpaces() int {
	ss := strings.Split(s.String(), "\n")
	return len(strings.Split(ss[len(ss)-1], " "))
}

func (s *str) Back(n int) {
	if n <= 0 {
		return
	}
	_s := s.String()
	s.Reset()
	s.WriteString(_s[:len(_s)-n])
}

func Fmt(code string) string {
	sb := &str{}
	lex := parser.NewV2Lexer(antlr.NewInputStream(code))
	tokens := lex.GetAllTokens()
	ident := 0
	var prevToken antlr.Token
	for _, cur := range tokens {
		cur := cur // Old golang may use the same variable for each iteration
		switch cur.GetTokenType() {
		case parser.V2LexerLBrack:

			if !sb.EndsWith("\n") && !sb.EndsWithSpaces() {
				sb.WriteString(" ")
				sb.WriteString("{\n")
			} else {
				makeIdent(sb, ident)
				sb.WriteString("{\n")
			}
			ident++
		case parser.V2LexerRBrack:
			if !sb.EndsWith("\n") && !sb.EndsWithSpaces() {
				sb.WriteString("\n")
				ident--
				continue
			}
			sb.Back(sb.CountSpaces())
			ident--
			sb.WriteString("\n")
			makeIdent(sb, ident)
			sb.WriteString("}\n")
			if ident == 0 {
				sb.WriteString("\n\n")
			}
		case parser.V2LexerNewLine:
			if !sb.EndsWith("\n") {
				sb.WriteString("\n")
			}
			makeIdent(sb, ident)
		case parser.V2LexerRParen:
			sb.WriteString(")")
		case parser.V2LexerRSquare:
			sb.WriteString("]")
		case parser.V2LexerLParen:
			if prevToken.GetTokenType() != parser.V2ParserIdentifier {
				sb.WriteString(" ")
			}
			sb.WriteString("(")
		case parser.V2LexerComment:
			if !sb.EndsWith("\n") && !sb.EndsWithSpaces() {
				sb.WriteString(" ")
			}
			sb.WriteString(cur.GetText())
		case parser.V2LexerElse:
			if sb.EndsWith("\n") {
				sb.Back(1)
			}
			sb.WriteString(" else")
		case parser.V2LexerSemi, parser.V2LexerComma, parser.V2LexerDot:
			sb.WriteString(cur.GetText())
		default:
			switch prevToken.GetTokenType() {
			case parser.V2LexerLSquare, parser.V2LexerLParen, parser.V2LexerDot:
			default:
				if !sb.EndsWith("\n") && !sb.EndsWithSpaces() {
					sb.WriteString(" ")
				}
			}
			if cur.GetTokenType() == parser.V2LexerFunction || cur.GetTokenType() == parser.V2LexerStruct {

				if prevToken != nil && !sb.EndsWith("\n\n") {
					switch prevToken.GetTokenType() {
					case parser.V2LexerAssign, parser.V2LexerComment:
					default:
						sb.WriteString("\n\n")
					}
				}

			}
			sb.WriteString(cur.GetText())
		}
		prevToken = cur
	}
	return strings.TrimSpace(sb.String())
}

func makeIdent(sb *str, n int) {
	if sb.EndsWithSpaces() && n == (sb.CountSpaces()/4) {
		return
	}
	for i := 0; i < n; i++ {
		sb.WriteString("    ")
	}
}
