package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestAssignStmt(t *testing.T) {
	code := `
x := [12, 18, 19]
x[0] := 19
println(x)
`
	expected := `[19 18 19]
`
	tester.GeneralTest(false, t, code, expected)
}

func TestMapAssignStmt(t *testing.T) {
	code := `
x := {1: 12, 2: 18, 3: 19}
x[1] := 18
println(x)
`
	expected := `map[1:18 2:18 3:19]
`
	tester.GeneralTest(false, t, code, expected)
}

func TestMapAssignStmt2(t *testing.T) {
	code := `
x := {}
x[1] := 18
println(x)
`
	expected := `map[1:18]
`
	tester.GeneralTest(false, t, code, expected)
}
