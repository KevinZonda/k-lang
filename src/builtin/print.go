package builtin

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/objType"
	"reflect"
)

func Print(v ...any) {
	for _, arg := range v {
		switch arg.(type) {
		case *objType.StructField:
			fmt.Println(arg.(*objType.StructField).Fields)
		default:
			fmt.Print(arg)
		}
	}
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

func Match(ns ...string) any {
	switch ns[0] {
	case "print":
		return Print
	case "println":
		return Println
	case "typeOf":
		return TypeOf
	}
	return nil
}

func Call(fn any, args []any) []any {
	if fn == nil {
		return nil // FIXME: Nil FUNC
	}
	var argsV []reflect.Value
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
	}
	vC := reflect.ValueOf(fn).Call(argsV)
	var vs []any
	for _, v := range vC {
		vs = append(vs, v.Interface())
	}

	return vs
}
