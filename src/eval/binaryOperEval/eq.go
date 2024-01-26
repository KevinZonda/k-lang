package binaryOperEval

import (
	"fmt"
	"math"
)

func Eq(left any, right any) bool {
	switch leftT := left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return EqFloat(leftT, right.(float64))
		case int:
			return EqFloat(leftT, float64(right.(int)))
		}
	case int:
		switch right.(type) {
		case float64:
			return EqFloat(float64(leftT), right.(float64))
		case int:
			return leftT == right.(int)
		}
	case bool:
		switch right.(type) {
		case bool:
			return leftT == right.(bool)
		}
	case string:
		return leftT == fmt.Sprint(right)
	case []any:
		if rightT, ok := right.([]any); ok {
			for i, v := range leftT {
				if !Eq(v, rightT[i]) {
					return false
				}
			}
			return true
		}
	}
	panic(fmt.Sprintf("cannot eq %T and %T", left, right))
}

const AllowedFloatDeviation = 0.00000000001

//                                    0.30000000000000004

func EqFloat(left float64, right float64) bool {
	return math.Abs(left-right) < AllowedFloatDeviation
}
