package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
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

func TestAssignTypeNum2(t *testing.T) {
	code := `
num x = "18"
`
	tester.ExpectPanic(t, code, func(exp string) bool {
		return strings.HasPrefix(exp, "TypeCheck Failed")
	})
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

func TestAssignTypeStruct(t *testing.T) {
	code := `
open "feat/staticType"

struct foo {
	int x
}
foo x = struct {x: 18}
println(x)
`
	tester.NoPanic(t, code)
}

func TestAssignTypeStruct2(t *testing.T) {
	code := `
open "feat/staticType"

struct foo {
	int x
    fn bar() {
		self.x = self.x + 1
	}
}
foo x = struct {x: 18}
x.bar()
print(x.x)
`
	tester.GeneralTest(false, t, code, "19")
}

func TestAssignTypeFunc3(t *testing.T) {
	code := `
fn f() -> int {
    return ""
}
f()
`
	tester.ExpectPanic(t, code, func(exp string) bool {
		return strings.HasPrefix(exp, "TypeCheck Failed")
	})
}

func TestAssignTypeFunc4(t *testing.T) {
	code := `
fn f() -> (int, string) {
    return 12.0, 12
}
x, z := f()
print(typeOf(x))
`
	tester.ExpectPanic(t, code, func(exp string) bool {
		return strings.HasPrefix(exp, "TypeCheck Failed")
	})
}

func TestAssignTypeFunc5(t *testing.T) {
	code := `
struct color {
	r: int
	g: int
	b: int
}
x := struct {
    r: 18,
	g: 18,
	b: 1
}
println(typeOf(x))
z := color(x)
print(typeOf(z))
`
	tester.GeneralTest(true, t, code, "struct\ncolor")
}
