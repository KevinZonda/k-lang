package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestUnaryExpr(t *testing.T) {
	code := `
x := 1
y := -x
println(y)
z := -y
println(z)
a := !true
println(a)
b := !false
println(b)
`
	expected := `-1
1
false
true
`
	tester.GeneralTest(false, t, code, expected)
}
