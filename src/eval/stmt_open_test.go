package eval_test

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestOpenStmt(t *testing.T) {
	fmt.Println(filepath.Abs("."))
	code := `
open "testFile/hello.k" as hello
`
	expected := "Hello from hello.k\n"
	generalTest(t, code, expected)
}

func TestMultiOpenStmt(t *testing.T) {
	fmt.Println(filepath.Abs("."))
	code := `
open "testFile/hello.k" as hello
open "testFile/hello.k" as z
open "testFile/hello.k" as m
`
	expected := "Hello from hello.k\n"
	generalTest(t, code, expected)
}
