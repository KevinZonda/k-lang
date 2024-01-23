package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type FuncBlock struct {
	Token   token.Token
	Name    *Identifier
	Args    []*FuncArg
	RetType *Type
	Body    *CodeBlock
}

func (f *FuncBlock) block() {}
func (f *FuncBlock) TokenValue() string {
	return f.Name.Value
}
func (f *FuncBlock) GetToken() token.Token {
	return f.Token
}

type FuncArg struct {
	Type *Type
	Name *Identifier
}

type CodeBlock struct {
	Token token.Token
	Nodes []Node
}

// TODO: Codeblock is a stmt, need stmt()?
func (c *CodeBlock) block() {}
func (c *CodeBlock) stmt()  {}
func (c *CodeBlock) TokenValue() string {
	return "CodeBlock"
}
func (c *CodeBlock) GetToken() token.Token {
	return c.Token
}
