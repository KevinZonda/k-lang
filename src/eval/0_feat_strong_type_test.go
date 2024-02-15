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

func TestAssignType2(t *testing.T) {
	code := `
open "feat/staticType"

struct foo {}
foo x
x = 11
println(x)
`
	tester.ExpectPanic(t, code, func(exp string) bool {
		return strings.HasPrefix(exp, "TypeCheck Failed")
	})
}

func TestAssignType3(t *testing.T) {
	code := `
open "feat/staticType"

struct foo {}
x = foo{}
x = 11
println(x)
`
	tester.ExpectPanic(t, code, func(exp string) bool {
		return strings.HasPrefix(exp, "TypeCheck Failed")
	})
}

func TestAssignTypeNum(t *testing.T) {
	code := `
open "feat/staticType"
num x = 18
x = 18.0
println(x)
`
	tester.NoPanic(t, code)
}

func TestAssignTypeLambda(t *testing.T) {
	code := `
open "feat/staticType"

map x = fn () {
    println(11)
}
println(typeOf(x))
x()
`
	tester.ExpectPanic(t, code, func(exp string) bool {
		return strings.HasPrefix(exp, "TypeCheck Failed")
	})
}
