package eval_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"path/filepath"
	"testing"
)

func TestOpenStmt(t *testing.T) {
	eval.ResetGlobal()
	fmt.Println(filepath.Abs("."))
	code := `
open "testFile/hello.k" as hello
`
	expected := "Hello from hello.k\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestMultiOpenStmt(t *testing.T) {
	eval.ResetGlobal()
	fmt.Println(filepath.Abs("."))
	code := `
open "testFile/hello.k" as hello
open "testFile/hello.k" as z
open "testFile/hello.k" as m
`
	expected := "Hello from hello.k\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestOpenLibStmt(t *testing.T) {
	eval.ResetGlobal()
	fmt.Println(filepath.Abs("."))
	code := `
open "std/string"

x := "This is a string"
typeOf(x)
`
	expected := ""
	tester.GeneralTest(true, t, code, expected)
}

func TestAllOpenFeat(t *testing.T) {
	eval.ResetGlobal()
	fmt.Println(filepath.Abs("."))
	code := `
open "feat/verbose:on"
open "feat/visualize:off"
open "feat/staticType"
open "feat/unknownVarNil"
x

`
	tester.NoPanic(t, code)
}
