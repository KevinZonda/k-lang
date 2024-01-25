package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

//func TestDotExpr(t *testing.T) {
//	code := `
//a := 11
//z := "ZZZ"
//a.z.y()
//`
//	a, _ := parserHelper.Ast(code)
//	eval.New(a, "").Do()
//}

func TestDotExpr2(t *testing.T) {
	code := `
open "testFile/open_func_call.k" as x

x.printHi()
`
	expected := `load open_func_call.k
Calling open_func_call.printHi()
Hi
`
	tester.GeneralTest(false, t, code, expected)
}

func TestDotExpr3(t *testing.T) {
	code := `
c := struct {
  R: 255,
  G: 114,
  B: 819
}

println(c.R)
println(c.G)
`
	expected := `255
114
`
	tester.GeneralTest(true, t, code, expected)
}

func TestDotExprStdLib(t *testing.T) {
	code := `
open "std/console"

console.writeln("HI!")
`
	expected := `HI!
`
	tester.GeneralTest(false, t, code, expected)
}

func TestDotExprStructLambda(t *testing.T) {
	code := `
struct M {
  x = 10
  M = fn(x) {
     println(self.x)
     println(x)
  }
}
M{}.M(20)
`
	expected := `10
20
`
	tester.GeneralTest(false, t, code, expected)
}
