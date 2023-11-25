package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestDotExpr(t *testing.T) {
	code := `
a := 11
z := "ZZZ"
a.z.y()
`
	a, _ := parserHelper.Ast(code)
	eval.New(a, "").Do()
}

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
c := colour {
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
