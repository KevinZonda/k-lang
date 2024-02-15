package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"strings"
	"testing"
)

func TestAssignType(t *testing.T) {
	code := `
open "feat/staticType"
int x = 18
x = "18"
println(x)
`
	tester.ExpectPanic(t, code, func(exp string) bool {
		return strings.HasPrefix(exp, "TypeCheck Failed")
	})
}
