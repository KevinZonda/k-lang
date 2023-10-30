package binaryOperEval

import "fmt"

func Mod(left any, right any) int {
	switch left.(type) {
	case int:
		switch right.(type) {
		case int:
			return left.(int) % right.(int)
		}
	}
	panic(fmt.Sprintf("cannot mod %T and %T", left, right))
}
