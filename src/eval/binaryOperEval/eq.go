package binaryOperEval

import "fmt"

func Eq(left any, right any) bool {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return left.(float64) == right.(float64)
		case int:
			return left.(float64) == float64(right.(int))
		}
	case int:
		switch right.(type) {
		case float64:
			return float64(left.(int)) == right.(float64)
		case int:
			return left.(int) == right.(int)
		}
	case bool:
		switch right.(type) {
		case bool:
			return left.(bool) == right.(bool)
		}
	}
	panic(fmt.Sprintf("cannot eq %T and %T", left, right))
}
