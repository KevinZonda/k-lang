package objType

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

type ILibrary interface {
	FuncCall(name string, args []any) any
}

type Object struct {
	Kind Kind
	Val  any
}

type Kind string

const (
	Lambda  Kind = "Lambda"
	Func    Kind = "Func"
	Value   Kind = "Val"
	EvalObj Kind = "Eval"
	Library Kind = "Library"
	//Array   Kind = "Array"
	//Map     Kind = "Map"
	Struct Kind = "Struct"
)

func (o *Object) TypeOf() string {
	return string(o.Kind)
}

func NewObj(k Kind, val any) *Object {
	return &Object{Kind: k, Val: val}
}

func (o *Object) ToStruct() *StructField {
	return o.Val.(*StructField)
}

func (o *Object) ToLambda() *node.LambdaExpr {
	return o.Val.(*node.LambdaExpr)
}

func (o *Object) ToFunc() *node.FuncBlock {
	return o.Val.(*node.FuncBlock)
}

func (o *Object) ToLib() ILibrary {
	return o.Val.(ILibrary)
}

func (o *Object) ToValue() any {
	return o.Val
}

func (o *Object) Is(kinds ...Kind) bool {
	for _, kind := range kinds {
		if o.Kind == kind {
			return true
		}
	}
	return false
}
