package binaryOperEval

import "fmt"

func Greater(left any, right any) bool {
	switch leftT := left.(type) {
	case float64:
		switch rightT := right.(type) {
		case float64:
			return leftT > rightT
		case int:
			return leftT > float64(rightT)
		}
	case int:
		switch rightT := right.(type) {
		case float64:
			return float64(leftT) > rightT
		case int:
			return leftT > rightT
		}
	}
	panic(fmt.Sprintf("cannot greater %T and %T", left, right))
}

func GreaterEq(left any, right any) bool {
	switch leftT := left.(type) {
	case float64:
		switch rightT := right.(type) {
		case float64:
			return leftT >= rightT
		case int:
			return leftT >= float64(rightT)
		}
	case int:
		switch rightT := right.(type) {
		case float64:
			return float64(leftT) >= rightT
		case int:
			return leftT >= rightT
		}
	}
	panic(fmt.Sprintf("cannot greater %T and %T", left, right))
}
