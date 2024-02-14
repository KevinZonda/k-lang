package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestBuiltinStringReplace(t *testing.T) {
	code := `
println("1122334".replace("11", "-0"))
`
	expected := `-022334
`
	tester.GeneralTest(false, t, code, expected)
}
