package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
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
