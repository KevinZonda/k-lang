package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type FuncBlock struct {
	Token   token.Token
	Name    *Identifier
	Args    []*FuncArg
	RetType *Identifier
	Body    *CodeBlock
}

func (f *FuncBlock) block() {}
func (f *FuncBlock) TokenValue() string {
	return f.Name.Value
}

type FuncArg struct {
	Type *Identifier
	Name *Identifier
}

type CodeBlock struct {
	Token token.Token
	Nodes []Node
}

func (c *CodeBlock) block() {}
func (c *CodeBlock) TokenValue() string {
	return "CodeBlock"
}
