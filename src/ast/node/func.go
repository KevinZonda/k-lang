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
	BinaryFx *reflect.Value
}

func cast(t *Type, v any) any {
	if t == nil {
		return v
	}
	if t.IsPlainType(TypeInt) {
		switch vT := v.(type) {
		case float64:
			return int(vT)
		case int:
			return vT
		default:
			panic("Not Possible Type Transformation")
		}
	}
	if t.IsPlainType(TypeNum) {
		switch vT := v.(type) {
		case float64:
			return vT
		case int:
			return float64(vT)
		default:

			panic("Not Possible Type Transformation")
		}
	}
	if t.IsPlainType(TypeString) {
		switch vT := v.(type) {
		case string:
			return vT
		default:
			panic("Not Possible Type Transformation")
		}
	}

	return v
}

func (f *FuncBlock) EvalBinary(args []any) []any {
	hasParam := false
	paramEmpty := false
	for _, arg := range f.Args {
		if arg.Param {
			hasParam = true
			paramEmpty = arg.ParamEmpty
			break
		}
	}
	if hasParam && len(args) < len(f.Args) ||
		!hasParam && len(args) != len(f.Args) {
		if !paramEmpty {
			panic("Invalid number of arguments")
		}
	}

	var argsV []reflect.Value
	isParam := false
	var param []any = nil
	for i, arg := range args {
		if !isParam && f.Args[i].Param {
			isParam = true
		}
		if isParam {
			param = append(param, arg)
			continue
		}

		arg = cast(f.Args[i].Type, arg)
		if arg == nil {
			argsV = append(argsV, reflect.Zero(reflect.TypeOf((*int)(nil)).Elem())) // trick to pass nil val or get ZeroValue Panic
			continue
		}
		argsV = append(argsV, reflect.ValueOf(arg))
	}
	if isParam {
		argsV = append(argsV, reflect.ValueOf(param))
	}

	if len(argsV) == 0 {
		if len(f.Args) > 0 {
			if f.Args[0].ParamEmpty {
				argsV = append(argsV, reflect.ValueOf([]any{}))
			} else {
				panic("INTERNAL ERR: PARAM EMPTY NOT ALLOWED")
			}
		}
	}

	vals := f.BinaryFx.Call(argsV)
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
	Type       *Type
	Name       *Identifier
	Ref        bool
	Param      bool
	ParamEmpty bool
}

func (a *FuncArg) String() string {
	name := "Anonymous"
	if a.Name != nil {
		name = a.Name.Value
	}
	if a.Type == nil {
		return name
	}
	return fmt.Sprintf("%s: %s", name, a.Type.Name)
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
