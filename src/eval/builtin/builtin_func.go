package builtin

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
)

func (b BuiltIn) Print(v ...any) {
	for _, arg := range v {
		switch arg.(type) {
		default:
			fmt.Fprint(b.StdOut, arg)
		}
	}
}

func (b BuiltIn) Println(v ...any) {
	for _, arg := range v {
		b.Print(arg)
	}
	fmt.Fprintln(b.StdOut)
}

func TypeOf(v any) string {
	switch v.(type) {
	case obj.ILibrary:
		return "lib"
	case int:
		return "int"
	case float64:
		return "num"
	case string:
		return "string"
	case bool:
		return "bool"
	case []any:
		return "array"
	case map[any]any:
		return "map"
	case *obj.Object:
		if v.(*obj.Object).Is(obj.Func, obj.Lambda) {
			return "fn"
		}
		vT := v.(*obj.Object)
		if v.(*obj.Object).Is(obj.Struct) {
			return TypeOf(vT.ToStruct())
		}
		if vT.Type != nil {
			return v.(*obj.Object).Type.CodeName()
		}
		return TypeOf(vT.Value())
	case *obj.StructField:
		return v.(*obj.StructField).TypeAs.Name
	}
	return reflect.TypeOf(v).String()
}

func Len(v any) int {
	if v == nil {
		return 0
	}
	switch v.(type) {
	case string:
		return len([]rune(v.(string)))
	case []any:
		return len(v.([]any))
	case map[any]any:
		return len(v.(map[any]any))
	}
	return 0
}

func Range(v any) []any {
	if v == nil {
		return nil
	}
	switch v.(type) {
	case int:
		vi := v.(int)
		if vi < 0 {
			return nil
		}
		vs := make([]any, vi)
		for i := 0; i < vi; i++ {
			vs[i] = i
		}
		return vs
	case int32:
		return Range(int(v.(int32)))
	case int64:
		return Range(int(v.(int64)))
	}
	return nil
}

func Panic(vs ...any) {
	panic(fmt.Sprint(vs...))
}
