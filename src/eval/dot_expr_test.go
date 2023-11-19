package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
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
	generalTest(false, t, code, expected)
}
