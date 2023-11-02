package builtin

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"reflect"
)

func Print(v ...any) {
	fmt.Print(v...)
}

func Println(v ...any) {
	fmt.Println(v...)
}

func Match(n *node.Variable) any {
	switch n.Value[0].Name.Value {
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
