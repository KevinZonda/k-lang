package binaryOperEval

import (
	"fmt"
	"math"
)

func Pow(left any, right any) any {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return math.Pow(left.(float64), right.(float64))
		case int:
			return math.Pow(left.(float64), float64(right.(int)))
		}
	case int:
		switch right.(type) {
		case float64:
			return math.Pow(float64(left.(int)), right.(float64))
		case int:
			return int(math.Pow(float64(left.(int)), float64(right.(int))))
		}
	}
	panic(fmt.Sprintf("cannot pow %T and %T", left, right))
}
