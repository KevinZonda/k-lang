package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestMultiAssignStmt(t *testing.T) {
	code := `
fn foo() {
    return 11, 12
}
x = foo()
println(x)
`
	expected := `[11 12]
`
	tester.GeneralTest(false, t, code, expected)
}

func TestMultiAssignStmt2(t *testing.T) {
	code := `
fn foo() {
    return 11, 12
}
x, y = foo()
println(x)
`
	expected := `11
`
	tester.GeneralTest(false, t, code, expected)
}
