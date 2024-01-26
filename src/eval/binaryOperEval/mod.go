package binaryOperEval

import "fmt"

func Mod(left any, right any) int {
	switch leftT := left.(type) {
	case int:
		switch rightT := right.(type) {
		case int:
			return leftT % rightT
		}
	}
	panic(fmt.Sprintf("cannot mod %T and %T", left, right))
}
