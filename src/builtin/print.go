package builtin

import (
	"fmt"
	"reflect"
)

func Print(v ...any) {
	fmt.Print(v...)
}

func Println(v ...any) {
	fmt.Println(v...)
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

func Call(fn any, args []any) []reflect.Value {
	if fn == nil {
		return nil // FIXME: Nil FUNC
	}
	var argsV []reflect.Value
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
	}
	return reflect.ValueOf(fn).Call(argsV)
}
