package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type LambdaExpr struct {
	Token   token.Token
	Args    []*FuncArg
	RetType *Identifier
	Body    *CodeBlock
}

func (f *LambdaExpr) expr() {}
func (f *LambdaExpr) TokenValue() string {
	return f.Token.Value
}

func (f *LambdaExpr) ToFunc(name string) *FuncBlock {
	return &FuncBlock{
		Token:   f.Token,
		Name:    &Identifier{Token: f.Token, Value: name},
		Args:    f.Args,
		RetType: f.RetType,
		Body:    f.Body,
	}
}