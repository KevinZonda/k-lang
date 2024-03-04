package tester

import (
	"fmt"
	"testing"
)

func Assert(t *testing.T, v1 any, v2 ...any) {
	for _, v := range v2 {
		if v1 != v {
			panic(fmt.Sprintf("Expected %v, got %v", v2, v1))
			return
		}
	}
}

func AssertArray(t *testing.T, v1 []any, v2 ...[]any) {
	for _, v := range v2 {
		for idx, e := range v {
			if v1[idx] != e {
				t.Errorf("Expected %v, got %v", v2, v1)
				return
			}
		}
	}
}
