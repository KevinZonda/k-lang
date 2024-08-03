package obj

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"strings"
)

type ILibrary interface {
	GetObj(name string) *Object
	GetFunc(name string) *node.FuncBlock
}

type BuiltInLibrary interface {
	ILibrary
	AllFuncs() map[string]*node.FuncBlock
	AllObjs() map[string]*Object
}

type ILibraryCall interface {
	HasValue() bool
	Value() any
}

type Object struct {
	Kind      Kind
	val       any
	Ref       *Object
	Type      *node.Type
	Immutable bool
}

func (o *Object) WithType(t *node.Type) *Object {
	o.Type = t
	return o
}

func (o *Object) AutoType() *Object {
	if o.Type != nil || o.Ref != nil {
		return o
	}
	if !o.Is(Value) {
		return o
	}
	// TODO:
	return o
}

func (o *Object) Value() any {
	if o.Ref != nil {
		return o.Ref.Value()
	}
	return o.val
}

func (o *Object) isLoopRef(v any) bool {
	vT, ok := v.(*Object)
	if !ok {
		return false
	}
	if vT == o {
		return true
	}
	if vT.Ref != nil {
		return o.isLoopRef(vT.Ref)
	}
	return false
}

func (o *Object) SetValue(v any) {
	if o.isLoopRef(v) { // prevent ref loop
		return
	}
	if o.Ref != nil {
		//o.Ref.val = v
		//o.Ref.autoKind()
		o.Ref.SetValue(v)
		return
	}
	o.val = v
	o.autoKind()
}

func (o *Object) autoKind() {
	switch o.val.(type) {
	case *node.FuncBlock:
		o.Kind = Func
	case *node.LambdaExpr:
		o.Kind = Lambda
	case *StructField:
		o.Kind = Struct
	}
	return
}

func (o *Object) NoRef() *Object {
	o.Ref = nil
	return o
}

func (o *Object) CreateRef() *Object {
	return &Object{
		Kind: o.Kind,
		Ref:  o,
	}
}

func (o *Object) GetType() *node.Type {
	return o.Type
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
	Struct    Kind = "Struct"
	StructDef Kind = "StructDef"
)

func (o *Object) String() string {
	if o == nil {
		return "<nil object>"
	}
	sb := strings.Builder{}
	sb.WriteString("Object{ Kind: ")
	sb.WriteString(string(o.Kind))
	sb.WriteString(", Val: ")
	sb.WriteString(fmt.Sprint(o.Value()))
	sb.WriteString(", Addr: ")
	sb.WriteString(fmt.Sprintf("%p", o))
	sb.WriteString(", Ref: ")
	if o.Ref != nil {
		sb.WriteString(fmt.Sprintf("%p", o.Ref))
	} else {
		sb.WriteString("nil")
	}
	sb.WriteString(" }")
	return sb.String()
}

func NewObj(k Kind, val any) *Object {
	return &Object{Kind: k, val: val}
}

func NewImmutableObj(k Kind, val any) *Object {
	return &Object{Kind: k, val: val, Immutable: true}
}

func (o *Object) ToStruct() *StructField {
	return o.Value().(*StructField)
}

func (o *Object) ToLambda() *node.LambdaExpr {
	return o.Value().(*node.LambdaExpr)
}

func (o *Object) ToFunc() *node.FuncBlock {
	return o.Value().(*node.FuncBlock)
}

func (o *Object) ToLib() ILibrary {
	return o.Value().(ILibrary)
}

func (o *Object) ToStructDef() *node.StructBlock {
	return o.Value().(*node.StructBlock)
}

func (o *Object) ToValue() any {
	return o.Value()
}

func (o *Object) Is(kinds ...Kind) bool {
	if o == nil {
		return false
	}
	for _, kind := range kinds {
		if o.Kind == kind {
			return true
		}
	}
	return false
}

func (o *Object) Visualize() string {
	switch o.Kind {
	case Value:
		switch o.Value().(type) {
		case *StructField:
			return TreeAnyW("struct", o, false).String()
		}
	}
	return o.String()
}

var _ IVisualize = &StructField{}
