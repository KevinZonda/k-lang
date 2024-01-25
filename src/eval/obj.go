package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func cons(a any) *obj.Object {
	switch a.(type) {
	//case Eval:
	//	e := a.(Eval)
	//	return NewEvalObject(&e)
	case *Eval:
		return obj.NewObj(obj.EvalObj, a)
	case obj.ILibrary:
		return obj.NewObj(obj.Library, a)
	case *obj.Object:
		return a.(*obj.Object)
	case *node.LambdaExpr:
		return obj.NewObj(obj.Lambda, a)
	case *node.FuncBlock:
		return obj.NewObj(obj.Func, a)
	//case []any:
	//	return NewArrayObject(a.([]any))
	//case map[any]any:
	//	return NewMapObject(a.(map[any]any))
	case *obj.StructField:
		return obj.NewObj(obj.Struct, a)
	case obj.StructField:
		sf := a.(obj.StructField)
		return obj.NewObj(obj.Struct, &sf)
	default:
		return obj.NewObj(obj.Value, a)
	}
}
