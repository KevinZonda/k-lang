package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestCodeBlock(t *testing.T) {
	code := `
{
	println("Hello")
}
`
	tester.GeneralTest(true, t, code, "Hello\n")
}

func TestMainFunc(t *testing.T) {
	code := `
fn main() {
    println("Hello")
}
`
	tester.GeneralTest(true, t, code, "Hello\n")
}

func TestPanic(t *testing.T) {
	code := `
fn main() {
    panic("Hello")
}
`
	tester.ExpectPanic(t, code, func(e string) bool {
		return e == "Hello"
	})
}

func TestEvalVisualize(t *testing.T) {
	code := `
open "feat/visualize"
x = struct {
x: 18
z: 9
}
println(visualize(x))
`
	tester.NoPanic(t, code)
}
