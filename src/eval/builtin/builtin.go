package builtin

import (
	"io"
	"os"
	"reflect"
)

type BuiltIn struct {
	StdIn  io.ReadCloser
	StdOut io.WriteCloser
	StdErr io.WriteCloser
}

func NewBuiltIn() *BuiltIn {
	return &BuiltIn{
		StdIn:  os.Stdin,
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}
}

func (b *BuiltIn) Call(fn any, args []any) []any {
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

func (b *BuiltIn) Match(ns ...string) any {
	switch ns[0] {
	case "print":
		return b.Print
	case "println":
		return b.Println
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
