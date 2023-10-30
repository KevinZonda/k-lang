package binaryOperEval

import "fmt"

func And(left any, right any) bool {
	switch left.(type) {
	case bool:
		switch right.(type) {
		case bool:
			return left.(bool) && right.(bool)
		}
	}
	panic(fmt.Sprintf("cannot and %T and %T", left, right))
}
