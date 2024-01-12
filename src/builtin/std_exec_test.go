package builtin_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestStdCmdLib(t *testing.T) {

	code := `
open "std/exec"
x = exec.cmd("echo", "hello")
println(x)
`
	expected := `hello

`

	tester.GeneralTest(false, t, code, expected)
}
