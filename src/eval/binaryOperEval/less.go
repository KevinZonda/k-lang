package binaryOperEval

import "fmt"

func Less(left any, right any) bool {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return left.(float64) < right.(float64)
		case int:
			return left.(float64) < float64(right.(int))
		}
	case int:
		switch right.(type) {
		case float64:
			return float64(left.(int)) < right.(float64)
		case int:
			return left.(int) < right.(int)
		}
	}
	panic(fmt.Sprintf("cannot less %T and %T", left, right))
}

func LessEq(left any, right any) bool {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return left.(float64) <= right.(float64)
		case int:
			return left.(float64) <= float64(right.(int))
		}
	case int:
		switch right.(type) {
		case float64:
			return float64(left.(int)) <= right.(float64)
		case int:
			return left.(int) <= right.(int)
		}
	}
	panic(fmt.Sprintf("cannot less %T and %T", left, right))
}
