package binaryOperEval

import "fmt"

func Div(left any, right any) float64 {
	switch leftT := left.(type) {
	case float64:
		switch rightT := right.(type) {
		case float64:
			return leftT / rightT
		case int:
			return leftT / float64(rightT)
		}
	case int:
		switch rightT := right.(type) {
		case float64:
			return float64(leftT) / rightT
		case int:
			return float64(leftT) / float64(rightT)
		}
	}
	panic(fmt.Sprintf("cannot div %T and %T", left, right))
}
