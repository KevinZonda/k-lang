package tester

import "testing"

func Assert(t *testing.T, v1 any, v2 ...any) {
	for _, v := range v2 {
		if v1 == v {
			return
		}
	}
	t.Errorf("Expected %v, got %v", v2, v1)
}
