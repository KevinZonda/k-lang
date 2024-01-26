package binaryOperEval

import "fmt"

func And(left any, right any) bool {
	switch leftT := left.(type) {
	case bool:
		switch rightT := right.(type) {
		case bool:
			return leftT && rightT
		}
	}
	panic(fmt.Sprintf("cannot and %T and %T", left, right))
}
