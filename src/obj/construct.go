package obj

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

func Construct(a any) *Object {
	switch a.(type) {
	case ILibrary:
		return NewObj(Library, a)
	case *Object:
		return a.(*Object)
	case *node.LambdaExpr:
		return NewObj(Lambda, a)
	case *node.FuncBlock:
		return NewObj(Func, a)
	case *node.StructBlock:
		return NewObj(StructDef, a)
	//case []any:
	//	return NewArrayObject(a.([]any))
	//case map[any]any:
	//	return NewMapObject(a.(map[any]any))
	case *StructField:
		return NewObj(Struct, a)
	case StructField:
		sf := a.(StructField)
		return NewObj(Struct, &sf)
	default:
		return NewObj(Value, a)
	}
}
