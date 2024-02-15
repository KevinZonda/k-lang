package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
)

func (e *Eval) TypeCheck(t *node.Type, v any) bool {
	if t == nil || !e.FeatStaticType {
		return true
	}
	if v == nil {
		return t.Nullable
	}
	switch v.(type) {
	case []any:
		return t.Array
	case map[any]any:
		return t.Map
	case int:
		return t.Name == "int"
	case float64:
		return t.Name == "num"
	case string:
		return t.Name == "string" || t.Name == "str"
	default:
		// TODO: more type check
		return true
	}
}

func (e *Eval) TypeCheckOrPanic(t *node.Type, v any) {
	checked := e.TypeCheck(t, v)
	if !checked {
		panic(fmt.Sprintf("TypeCheck Failed, expected %s, got %s (%s)", t.CodeName(), reflect.TypeOf(v), v))
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
	default:
		// TODO: more type check
		return nil
	}
}
