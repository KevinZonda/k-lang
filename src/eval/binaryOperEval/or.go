package binaryOperEval

import "fmt"

func Or(left any, right any) bool {
	switch left.(type) {
	case bool:
		switch right.(type) {
		case bool:
			return left.(bool) || right.(bool)
		}
	}
	panic(fmt.Sprintf("cannot or %T and %T", left, right))
}
