package binaryOperEval

import (
	"fmt"
)

func Sub(left any, right any) any {
	switch leftT := left.(type) {
	case float64:
		switch rightT := right.(type) {
		case float64:
			return leftT - rightT
		case int:
			return leftT - float64(rightT)
		}
	case int:
		switch rightT := right.(type) {
		case float64:
			return float64(leftT) - rightT
		case int:
			return leftT - rightT
		}
	}
	panic(fmt.Sprintf("cannot minus %T and %T", left, right))
}
