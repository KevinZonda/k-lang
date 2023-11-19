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

func Match(ns ...string) any {
	switch ns[0] {
	case "print":
		return Print
	case "println":
		return Println
	}
	return nil
}

func Call(fn any, args []any) []reflect.Value {
	if fn == nil {
		return nil // FIXME: Nill FUNC
	}
	var argsV []reflect.Value
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
	}
	return reflect.ValueOf(fn).Call(argsV)
}
