package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type Identifier struct {
	baseExpr
	Token token.Token
	Value string
}

func (i *Identifier) TokenValue() string {
	return i.Value
}
