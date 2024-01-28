package stringProcess

type Token struct {
	Kind  TokenKind
	Value string
}

func Parse(mode Mode, s string, c uint8) []Token {
	lex := &Lexer{
		raw:  []rune(s),
		mode: mode,
		sep:  c,
	}
	var parts []Token
	for lex.HasNext() {
		kind, token := lex.Next()
		parts = append(parts, Token{kind, token})
	}
	return parts
}
