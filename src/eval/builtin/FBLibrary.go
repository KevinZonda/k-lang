package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
)

type FBLibrary struct {
	F     map[string]*node.FuncBlock
	P     map[string]*obj.Object
	IO    obj.StdIO
	IODep bool
}

func (f FBLibrary) IsIODep() bool {
	return f.IODep
}

func (f FBLibrary) GetFunc(name string) *node.FuncBlock {
	if fx, ok := f.F[name]; ok {
		return fx
	}
	return nil
}

func (f FBLibrary) GetObj(name string) *obj.Object {
	if len(f.P) == 0 {
		return nil
	}
	if o, ok := f.P[name]; ok {
		return o
	}
	return nil

}

//func (f FBLibrary) FuncCall(name string, args []any) obj.ILibraryCall {
//	if f, ok := f.V[name]; ok {
//		v := f.EvalBinary(args)
//		switch len(v) {
//		case 0:
//			return resultNoVal()
//		case 1:
//			return resultVal(v[0])
//		default:
//			return resultVal(v)
//		}
//	}
//	panic("Function not found")
//}

func (f FBLibrary) GetObjList() map[string]*obj.Object {
	return nil
}

func NewTypeN(name string) *node.Type {
	return &node.Type{Name: name}
}

func FxToFuncBlock(f any) *node.FuncBlock {
	return &node.FuncBlock{
		Name:     node.NewIdentifier(""),
		Args:     fxToArgs(f),
		RetType:  fxToRetT(f),
		BinaryFx: f,
	}
}

func fxToArgs(fx any) []*node.FuncArg {
	f := reflect.TypeOf(fx)
	var args []*node.FuncArg
	n := f.NumIn()
	for i := 0; i < n; i++ {
		inT := f.In(i)
		args = append(args, &node.FuncArg{
			Name: node.NewIdentifier(""),
			Type: argT(inT),
		})
	}
	return args
}

func fxToRetT(fx any) []*node.Type {
	f := reflect.TypeOf(fx)
	var args []*node.Type
	n := f.NumOut()
	for i := 0; i < n; i++ {
		inT := f.Out(i)
		args = append(args, argT(inT))
	}
	return args
}

func argT(fx reflect.Type) *node.Type {
	switch fx.Kind() {
	case reflect.Int:
		return NewTypeN(node.TypeInt)
	case reflect.String:
		return NewTypeN(node.TypeString)
	case reflect.Bool:
		return NewTypeN(node.TypeBool)
	case reflect.Float64:
		return NewTypeN(node.TypeNum)
	case reflect.Interface:
		return NewTypeN(node.TypeAny)
	case reflect.Slice:
		switch fx.Elem().Kind() {
		case reflect.Interface:
			t := NewTypeN("")
			t.Array = true
			return t
		default:
			panic("Invalid elem type: " + fx.Elem().String())
		}
	}
	panic("Invalid type: " + fx.String())
}

var _ obj.ILibrary = &FBLibrary{}
