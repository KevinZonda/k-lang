package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/ast/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenValue() string {
	return i.Value
}
