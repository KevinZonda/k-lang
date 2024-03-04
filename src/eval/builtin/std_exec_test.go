package builtin_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
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

func TestRange(t *testing.T) {

	code := `
for (i : range(5)) {
    println(i)
}
`
	expected := `0
1
2
3
4
`

	tester.GeneralTest(false, t, code, expected)
}
