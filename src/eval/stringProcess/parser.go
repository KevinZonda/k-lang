package stringProcess

type Token struct {
	Kind  TokenKind
	Value string
}

func Parse(mode Mode, s string) []Token {
	lex := &Lexer{
		raw:  []rune(s),
		mode: mode,
	}
	var parts []Token
	for lex.HasNext() {
		kind, token := lex.Next()
		parts = append(parts, Token{kind, token})
	}
	return parts
}
