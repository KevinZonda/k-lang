package node

import "git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenValue() string {
	return i.Value
}
