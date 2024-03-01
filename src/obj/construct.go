package obj

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func Construct(a any) *Object {
	switch aT := a.(type) {
	case ILibrary:
		return NewObj(Library, a)
	case *Object:
		return aT
	case *node.LambdaExpr:
		t := node.NewType()
		t.Func = true
		return NewObj(Lambda, a).WithType(t)
	case *node.FuncBlock:
		t := node.NewType()
		t.Func = true
		return NewObj(Func, a).WithType(t)
	case *node.StructBlock:
		return NewObj(StructDef, a).WithType(node.NewType().WithName(aT.Name))
	//case []any:
	//	return NewArrayObject(a.([]any))
	//case map[any]any:
	//	return NewMapObject(a.(map[any]any))
	case *StructField:
		t := aT.TypeAs
		if t == nil {
			t = node.NewType()
		}
		return NewObj(Struct, a).WithType(t)
	case StructField:
		t := aT.TypeAs
		if t == nil {
			t = node.NewType()
		}
		return NewObj(Struct, &aT).WithType(t)
	default:
		o := NewObj(Value, a)
		switch a.(type) {
		case int:
			o.WithType(node.NewType().WithName(node.TypeInt))
		case float64:
			o.WithType(node.NewType().WithName(node.TypeNum))
		case string:
			o.WithType(node.NewType().WithName(node.TypeString))
		case bool:
			o.WithType(node.NewType().WithName(node.TypeBool))
		case []any:
			t := node.NewType()
			t.Array = true
			o.WithType(t)
		case map[any]any:
			t := node.NewType()
			t.Map = true
			o.WithType(t)
		}
		return o
	}
}
