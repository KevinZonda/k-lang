package eval_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
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
	GeneralTest(false, t, code, expected)
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
	GeneralTest(false, t, code, expected)
}
