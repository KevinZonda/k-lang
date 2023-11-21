package binaryOperEval

import (
	"fmt"
)

func Add(left any, right any) any {
	// string + any
	if r, ok := strAddAny(left, right); ok {
		return r
	}

	// float + any
	if r, ok := floatAddAny(left, right); ok {
		return r
	}

	// any + float
	if r, ok := floatAddAny(right, left); ok {
		return r
	}

	// int + any
	if r, ok := intAddAny(left, right); ok {
		return r
	}

	// any + int
	if r, ok := intAddAny(right, left); ok {
		return r
	}

	// arr + arr
	if r, ok := arrAndArr(left, right); ok {
		return r
	}
	panic(fmt.Sprintf("cannot add %T and %T", left, right))
}

func arrAndArr(left, right any) (any, bool) {
	switch left.(type) {
	case []any:
		switch right.(type) {
		case []any:
			return append(left.([]any), right.([]any)...), true
		default:
			return append(left.([]any), right), true
		}
	}
	return nil, false
}

func strAddAny(left, right any) (string, bool) {
	switch left.(type) {
	case string:
		switch right.(type) {
		case string:
			return left.(string) + right.(string), true
		case float64:
			return left.(string) + fmt.Sprintf("%f", right.(float64)), true
		case int:
			return left.(string) + fmt.Sprintf("%d", right.(int)), true
		}
	}
	switch right.(type) {
	case string:
		switch left.(type) {
		case string:
			return left.(string) + right.(string), true
		case float64:
			return fmt.Sprintf("%f", left.(float64)) + right.(string), true
		case int:
			return fmt.Sprintf("%d", left.(int)) + right.(string), true
		}

	}
	return "", false
}

func floatAddAny(left, right any) (float64, bool) {
	switch left.(type) {
	case float64:
		switch right.(type) {
		case float64:
			return left.(float64) + right.(float64), true
		case int:
			return left.(float64) + float64(right.(int)), true
		}
	}
	return 0, false
}

func intAddAny(left, right any) (any, bool) {
	switch left.(type) {
	case int:
		switch right.(type) {
		case float64:
			return float64(left.(int)) + right.(float64), true
		case int:
			return left.(int) + right.(int), true
		}
	}
	return 0, false
}
