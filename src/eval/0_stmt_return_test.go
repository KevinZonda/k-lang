package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestReturn(t *testing.T) {
	code := `
fn foo() {
	return 12
}
println(foo())
`
	expected := `12
`
	tester.GeneralTest(false, t, code, expected)
}
