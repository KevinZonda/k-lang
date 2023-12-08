package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestArrayIndexExpr(t *testing.T) {
	code := `
x = [18, 19, 20]
println(x[1])
`
	expected := `19
`
	tester.GeneralTest(true, t, code, expected)
}

func TestMapIndexExpr(t *testing.T) {
	code := `
x = {"a": 19, "b": 20}
println(x["a"])
println(x["b"])
println(x["c"])
`
	expected := `19
20

`
	tester.GeneralTest(true, t, code, expected)
}
