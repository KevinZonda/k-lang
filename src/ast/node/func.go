package node

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"reflect"
)

type FuncBlock struct {
	Token    token.Token
	Name     *Identifier
	Args     []*FuncArg
	RetType  []*Type
	Body     *CodeBlock
	BinaryFx any
}

func (f *FuncBlock) EvalBinary(args []any) []any {
	argsV := make([]reflect.Value, len(f.Args))
	for i, arg := range args {
		if arg == nil {
			argsV[i] = reflect.Zero(reflect.TypeOf((*int)(nil)).Elem()) // trick to pass nil val or get ZeroValue Panic
			continue
		}
		argsV[i] = reflect.ValueOf(arg)
	}
	vals := reflect.ValueOf(f.BinaryFx).Call(argsV)
	switch len(vals) {
	case 0:
		return nil
	case 1:
		return []any{vals[0].Interface()}
	default:
		var rst []any
		for _, v := range vals {
			rst = append(rst, v.Interface())
		}
		return rst
	}
}

func (f *FuncBlock) block() {}
func (f *FuncBlock) TokenValue() string {
	return f.Name.Value
}
func (f *FuncBlock) GetToken() token.Token {
	return f.Token
}
func (f *FuncBlock) String() string {
	name := "Anonymous"
	if f.Name != nil {
		name = f.Name.Value
	}
	return fmt.Sprintf("Function<%s@%p>", name, f)
}

type FuncArg struct {
	Type *Type
	Name *Identifier
	Ref  bool
}

type CodeBlock struct {
	Token token.Token
	Nodes []Node
}

func (c *CodeBlock) block() {}
func (c *CodeBlock) stmt()  {}
func (c *CodeBlock) TokenValue() string {
	return "CodeBlock"
}
func (c *CodeBlock) GetToken() token.Token {
	return c.Token
}
