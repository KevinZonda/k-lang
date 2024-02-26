package eval

import (
	"fmt"
)

func asType[T any](v any) T {
	if vT, ok := v.(T); ok {
		return vT
	}
	var r T
	panic(fmt.Sprintf("expect type: %T, but got: %T", r, v))
}

func possibleType[T any](v any) (T, bool) {
	vT, ok := v.(T)
	return vT, ok
}

func anyArr[T any](v []T) []any {
	var arr []any
	for _, a := range v {
		arr = append(arr, a)
	}
	return arr
}
