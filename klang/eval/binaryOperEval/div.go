package binaryOperEval

import "fmt"

func Div(left any, right any) float64 {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return left.(float64) / right.(float64)
		case int:
			return left.(float64) / float64(right.(int))
		}
	case int:
		switch right.(type) {
		case float64:
			return float64(left.(int)) / right.(float64)
		case int:
			return float64(left.(int)) / float64(right.(int))
		}
	}
	panic(fmt.Sprintf("cannot div %T and %T", left, right))
}
