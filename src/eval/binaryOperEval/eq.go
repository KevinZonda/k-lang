package binaryOperEval

import (
	"fmt"
	"math"
)

func Eq(left any, right any) bool {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return EqFloat(left.(float64), right.(float64))
		case int:
			return EqFloat(left.(float64), float64(right.(int)))
		}
	case int:
		switch right.(type) {
		case float64:
			return EqFloat(float64(left.(int)), right.(float64))
		case int:
			return left.(int) == right.(int)
		}
	case bool:
		switch right.(type) {
		case bool:
			return left.(bool) == right.(bool)
		}
	case string:
		return left.(string) == fmt.Sprint(right)
	}
	panic(fmt.Sprintf("cannot eq %T and %T", left, right))
}

const AllowedFloatDeviation = 0.00000000001

//                                    0.30000000000000004

func EqFloat(left float64, right float64) bool {
	return math.Abs(left-right) < AllowedFloatDeviation
}
