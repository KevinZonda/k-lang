package stringProcess

import (
	"strings"
)

type Lexer struct {
	raw  []rune
	pos  int
	mode Mode
}

func (l *Lexer) SetRaw(raw string) {
	l.raw = []rune(raw)
	l.pos = 0
	l.mode = ModeNormal
}

func (l *Lexer) HasNext() bool {
	return l.pos < len(l.raw)
}

func (l *Lexer) NextToken() string {
	if l.pos >= len(l.raw) {
		return ""
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
			if !hasMap {
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
	return buf.String()
}
