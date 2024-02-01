package builtin

import (
	"fmt"
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
	if k, ok := v.(ITypeOf); ok {
		return k.TypeOf()
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
