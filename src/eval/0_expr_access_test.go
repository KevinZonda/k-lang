package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestAccessByIndex(t *testing.T) {
	code := `
x := [12, 18, 19]
println(x[0])
`
	expected := `12
`
	tester.GeneralTest(false, t, code, expected)
}

func TestMapAccessByIndex(t *testing.T) {
	code := `
x := {"a": 12, "z": 18, "n": 19}
println(x["a"])
`
	expected := `12
`
	tester.GeneralTest(false, t, code, expected)
}
