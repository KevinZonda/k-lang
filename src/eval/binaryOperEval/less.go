package binaryOperEval

import "fmt"

func Less(left any, right any) bool {
	switch leftT := left.(type) {
	case float64:
		switch rightT := right.(type) {
		case float64:
			return leftT < rightT
		case int:
			return leftT < float64(rightT)
		}
	case int:
		switch rightT := right.(type) {
		case float64:
			return float64(leftT) < rightT
		case int:
			return leftT < rightT
		}
	}
	panic(fmt.Sprintf("cannot less %T and %T", left, right))
}

func LessEq(left any, right any) bool {
	switch leftT := left.(type) {
	case float64:
		switch rightT := right.(type) {
		case float64:
			return leftT <= rightT
		case int:
			return leftT <= float64(rightT)
		}
	case int:
		switch rightT := right.(type) {
		case float64:
			return float64(leftT) <= rightT
		case int:
			return leftT <= rightT
		}
	}
	panic(fmt.Sprintf("cannot less %T and %T", left, right))
}
