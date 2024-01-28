package binaryOperEval

import (
	"fmt"
	"math"
)

func Pow(left any, right any) any {
	switch leftT := left.(type) {
	case float64:
		switch rightT := right.(type) {
		case float64:
			return math.Pow(leftT, rightT)
		case int:
			return math.Pow(leftT, float64(rightT))
		}
	case int:
		switch rightT := right.(type) {
		case float64:
			return math.Pow(float64(leftT), rightT)
		case int:
			return int(math.Pow(float64(leftT), float64(rightT)))
		}
	}
	panic(fmt.Sprintf("cannot pow %T and %T", left, right))
}
