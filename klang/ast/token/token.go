package token

type Kind int

type Token struct {
	Kind    Kind
	Literal string
}
