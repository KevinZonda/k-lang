package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestNum(t *testing.T) {
	code := `x := .2
print(x)`
	expected := `0.2`
	tester.GeneralTest(false, t, code, expected)
}

func TestArrayInitializer(t *testing.T) {
	code := `x := [1, 2, 3, 4, 5]
print(x)`
	expected := `[1 2 3 4 5]`
	tester.GeneralTest(false, t, code, expected)
}

func TestMapInitializer(t *testing.T) {
	code := `x := map{
    1 : "one",
    13: "Thirteen"
}
print(x)`
	expected := `map[1:one 13:Thirteen]`
	tester.GeneralTest(true, t, code, expected)
}

func TestZeroValueInt(t *testing.T) {
	code := `
struct M{ int x }
print(M{}.x)
`
	expected := `0`
	tester.GeneralTest(true, t, code, expected)
}

func TestZeroValueFloat(t *testing.T) {
	code := `
struct M{ num x }
print(M{}.x)
`
	expected := `0`
	tester.GeneralTest(true, t, code, expected)
}

func TestZeroValueString(t *testing.T) {
	code := `
struct M{ str x }
print(M{}.x)
`
	expected := ``
	tester.GeneralTest(true, t, code, expected)
}

func TestZeroValueBool(t *testing.T) {
	code := `
struct M{ bool x }
print(M{}.x)
`
	expected := `false`
	tester.GeneralTest(true, t, code, expected)
}

func TestFloat(t *testing.T) {
	code := `
x := 1.2
print(x)
`
	expected := `1.2`
	tester.GeneralTest(true, t, code, expected)
}
