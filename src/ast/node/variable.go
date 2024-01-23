package node

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"strings"
)

type Variable struct {
	Token token.Token
	Value []*BaseVariable
}

func (v *Variable) expr() {}

func (v *Variable) TokenValue() string {
	var n []string
	for _, val := range v.Value {
		n = append(n, val.String())
	}
	return strings.Join(n, ".")
}

func (v *Variable) GetToken() token.Token {
	return v.Token
}

type BaseVariable struct {
	Token token.Token
	Name  *Identifier
	Index []Expr
}

func (b *BaseVariable) expr() {}

func (v *BaseVariable) String() string {
	if len(v.Index) > 0 {
		return v.Name.Value + "[" + v.Index[0].TokenValue() + "]"
	}
	return v.Name.Value
}

func (v *BaseVariable) TokenValue() string {
	return "BaseVariable"
}

func (v *BaseVariable) GetToken() token.Token {
	return v.Token
}
