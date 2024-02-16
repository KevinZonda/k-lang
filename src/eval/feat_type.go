package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
)

func (e *Eval) TypeCheck(t *node.Type, v any) bool {
	if t == nil {
		return true
	}
	if v == nil {
		return t.Nullable
	}
	switch v.(type) {
	case []any:
		return t.Package == "" && t.Array
	case map[any]any:
		return t.Package == "" && t.Map
	case int:
		return t.Package == "" && (t.Name == "int" || t.Name == "num")
	case float64:
		return t.Package == "" && t.Name == "num"
	case string:
		return t.Package == "" && (t.Name == "string" || t.Name == "str")
	case *node.LambdaExpr, *node.FuncBlock:
		return t.Func
	case *obj.StructField:
		vT := v.(*obj.StructField)
		if vT.TypeAs == nil {
			return true
		}
		return vT.TypeAs.Name == t.Name
	case *obj.Object:
		vT := v.(*obj.Object)
		if vT.Is(obj.Func, obj.Lambda) {
			return t.Func
		}
		if vT.Is(obj.Value) {
			return e.TypeCheck(t, vT.Value())
		}
		if vT.Is(obj.Struct) {
			vStruct := vT.ToStruct()
			return e.TypeCheck(t, vStruct)
		}
		return true

	default:
		// TODO: more type check
		return true
	}
}

func (e *Eval) NormaliseWithType(t *node.Type, v any) any {
	if t == nil {
		return v
	}
	if t.Name == "num" && t.Package == "" {
		switch vT := v.(type) {
		case float64:
			return vT
		case int:
			return float64(vT)
		default:
			return v
			//panic("Not Possible Type Transformation")
		}
	}
	return v
}

func (e *Eval) TypeCheckOrPanic(t *node.Type, v any) {
	checked := e.TypeCheck(t, v)
	if !checked {
		panic(fmt.Sprintf("TypeCheck Failed, expected %s, got %s (%s)", t.CodeName(), reflect.TypeOf(v), fmt.Sprint(v)))
	}
}

func (e *Eval) AutoType(v any) *node.Type {
	switch v.(type) {
	case []any:
		return &node.Type{
			Array: true,
		}
	case map[any]any:
		return &node.Type{
			Map: true,
		}
	case *node.LambdaExpr, *node.FuncBlock:
		return &node.Type{
			Func: true,
		}
	case int:
		return &node.Type{
			Name: "int",
		}
	case float64:
		return &node.Type{
			Name: "num",
		}
	case string:
		return &node.Type{
			Name: "string",
		}
	case *obj.StructField:
		vT := v.(*obj.StructField)
		return vT.TypeAs
	case *obj.Object:
		vT := v.(*obj.Object)
		if vT.Is(obj.Lambda, obj.Func) {
			return &node.Type{
				Func: true,
			}
		}
		if vT.Is(obj.Struct) {
			return e.AutoType(vT.ToStruct())
		}
	}
	// TODO: more type check
	// fmt.Println(reflect.TypeOf(v))
	return nil
}
