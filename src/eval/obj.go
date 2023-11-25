package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/objType"
)

func cons(a any) *objType.Object {
	switch a.(type) {
	//case Eval:
	//	e := a.(Eval)
	//	return NewEvalObject(&e)
	case *Eval:
		return objType.NewObj(objType.EvalObj, a)
	case objType.ILibrary:
		return objType.NewObj(objType.Library, a)
	case *objType.Object:
		return a.(*objType.Object)
	case *node.LambdaExpr:
		return objType.NewObj(objType.Lambda, a.(*node.LambdaExpr))
	case *node.FuncBlock:
		return objType.NewObj(objType.Func, a.(*node.FuncBlock))
	//case []any:
	//	return NewArrayObject(a.([]any))
	//case map[any]any:
	//	return NewMapObject(a.(map[any]any))
	case *objType.StructField:
		return objType.NewObj(objType.Struct, a.(*objType.StructField))
	case objType.StructField:
		sf := a.(objType.StructField)
		return objType.NewObj(objType.Struct, &sf)
	default:
		return objType.NewObj(objType.Value, a)
	}
}
