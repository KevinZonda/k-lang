package builtin

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"reflect"
)

func Print(v ...any) {
	for _, arg := range v {
		switch arg.(type) {
		case *obj.StructField:
			printStruct(arg.(*obj.StructField))
		default:
			fmt.Print(arg)
		}
	}
}

func printStruct(v *obj.StructField) {
	if v == nil {
		return
	}
	if v.TypeAs != nil {
		fmt.Print(v.TypeAs.Name, " ")
	} else {
		fmt.Print("struct ")
	}
	fmt.Print("{")
	count := 0
	totalLen := v.Fields.Len()
	for pair := v.Fields.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Print(pair.Key, ": ")
		switch pair.Value.(type) {
		case *obj.StructField:
			printStruct(pair.Value.(*obj.StructField))
		default:
			fmt.Print(pair.Value)
		}
		if count < totalLen-1 {
			fmt.Print(", ")
		}
		count++
	}
	fmt.Print("}")
}

func Println(v ...any) {
	for _, arg := range v {
		Print(arg)
	}
	fmt.Println()
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

