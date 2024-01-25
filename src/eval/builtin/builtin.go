package builtin

import "reflect"

func Call(fn any, args []any) []any {
	if fn == nil {
		return nil // FIXME: Nil FUNC
	}
	var argsV []reflect.Value
	for _, arg := range args {
		val := reflect.ValueOf(arg)
		if val.IsValid() {
			// FIXME: Null Type Will Cause FunCall Panic!
			argsV = append(argsV, val)
		}
	}
	vC := reflect.ValueOf(fn).Call(argsV)
	var vs []any
	for _, v := range vC {
		vs = append(vs, v.Interface())
	}

	return vs
}

func Match(ns ...string) any {
	switch ns[0] {
	case "print":
		return Print
	case "println":
		return Println
	case "typeOf":
		return TypeOf
	case "panic":
		return Panic
	case "range":
		return Range
	case "len":
		return Len
	}
	return nil
}
