package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/builtin"
)

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
)

func (o *Object) TypeOf() string {
	return string(o.Kind)
}

func NewLambdaObject(val *node.LambdaExpr) *Object {
	return &Object{Kind: Lambda, Val: val}
}

func NewFuncObject(val *node.FuncBlock) *Object {
	return &Object{Kind: Func, Val: val}
}

func NewValueObject(val any) *Object {
	return &Object{Kind: Value, Val: val}
}

func NewEvalObject(val any) *Object {
	return &Object{Kind: EvalObj, Val: val}
}

func (o *Object) IsLambda() bool {
	return o.Kind == Lambda
}

func (o *Object) IsFunc() bool {
	return o.Kind == Func
}

func (o *Object) IsValue() bool {
	return o.Kind == Value
}

func (o *Object) IsEval() bool {
	return o.Kind == EvalObj
}

func (o *Object) IsLib() bool {
	return o.Kind == Library
}

func (o *Object) ToLambda() *node.LambdaExpr {
	return o.Val.(*node.LambdaExpr)
}

func (o *Object) ToFunc() *node.FuncBlock {
	return o.Val.(*node.FuncBlock)
}

func (o *Object) ToLib() builtin.ILibrary {
	return o.Val.(builtin.ILibrary)
}

func (o *Object) ToValue() any {
	return o.Val
}

func cons(a any) *Object {
	switch a.(type) {
	//case Eval:
	//	e := a.(Eval)
	//	return NewEvalObject(&e)
	case *Eval:
		return NewEvalObject(a.(*Eval))
	case builtin.ILibrary:
		return &Object{Kind: Library, Val: a}
	case *Object:
		return a.(*Object)
	case *node.LambdaExpr:
		return NewLambdaObject(a.(*node.LambdaExpr))
	case *node.FuncBlock:
		return NewFuncObject(a.(*node.FuncBlock))
	default:
		return NewValueObject(a)
	}
}
