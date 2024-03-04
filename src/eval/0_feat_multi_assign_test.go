package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
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
	code1 := `
fn foo() {
    return 11, 12
}
x, y = foo()
println(x)
`
	code2 := `
fn foo() {
    return [11, 12]
}
x, y = foo()
println(x)
`
	expected := `11
`
	tester.BatchRunSplit(t, tester.Strings(code1, code2), tester.Strings(expected, expected))
}
