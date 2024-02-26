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
