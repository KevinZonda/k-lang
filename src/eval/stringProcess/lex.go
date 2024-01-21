package stringProcess

import (
	"strings"
)

type Lexer struct {
	raw      []rune
	pos      int
	mode     Mode
	nextKind TokenKind
}

func (l *Lexer) SetRaw(raw string) {
	l.raw = []rune(raw)
	l.pos = 0
	l.mode = ModeNormal
}

func (l *Lexer) SetMode(mode Mode) {
	l.mode = mode
}

func (l *Lexer) HasNext() bool {
	return l.pos < len(l.raw)
}

func (l *Lexer) Next() (TokenKind, string) {
	if l.nextKind == KindVar {
		return l.nextVarToken()
	}
	return l.nextTextToken()
}

func (l *Lexer) nextVarToken() (TokenKind, string) {
	if l.pos >= len(l.raw) {
		return KindEOF, ""
	}
	buf := strings.Builder{}
	l.pos++ // ignore '{'
	for l.pos < len(l.raw) {
		c := l.raw[l.pos]
		if c == '}' {
			l.nextKind = KindText
			l.pos++ // ignore '}'
			return KindVar, buf.String()
		}
		buf.WriteRune(c)
		l.pos++
	}
	return KindVar, buf.String()
}

func (l *Lexer) nextTextToken() (TokenKind, string) {
	if l.pos >= len(l.raw) {
		return KindEOF, ""
	}
	buf := strings.Builder{}
	prevChar := rune(-1)
	for l.pos < len(l.raw) {
		c := l.raw[l.pos]
		needWriteCurrent := false
		if l.pos == len(l.raw)-1 {
			needWriteCurrent = true
		}

		if prevChar != rune(-1) {
			mapVal, hasMap := hasMapToken(l.mode, string(prevChar)+string(c))
			// fmt.Println("  -> MapVal: ", string(prevChar), string(c), mapVal, hasMap)
			if !hasMap {
				if prevChar == '{' && l.mode == ModeVar {
					l.nextKind = KindVar
					l.pos--
					//fmt.Println("Var: ", buf.String())
					//fmt.Println("Pos: ", l.pos)
					//fmt.Println("Cur: ", string(l.raw[l.pos]))
					return KindText, buf.String()
				}

				buf.WriteRune(prevChar)

				prevChar = c
				l.pos++
				goto endLoop
			}
			buf.WriteString(mapVal)

			prevChar = rune(-1)
			needWriteCurrent = false
			l.pos++
			goto endLoop
		} else {
			prevChar = c
			l.pos++
			goto endLoop
		}

	endLoop:
		if needWriteCurrent {
			buf.WriteRune(c)
		}
	}
	return KindText, buf.String()
}

type TokenKind int

const (
	KindEOF TokenKind = iota - 1
	KindText
	KindVar
)

func (k TokenKind) String() string {
	switch k {
	case KindEOF:
		return "EOF"
	case KindText:
		return "Text"
	case KindVar:
		return "Var"
	default:
		return "Unknown"
	}
}
