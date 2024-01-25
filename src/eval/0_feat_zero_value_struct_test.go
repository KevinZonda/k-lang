package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestZeroValueUsable(t *testing.T) {
	code := `
struct Ptr {
  int x
  int y
}

p := Ptr{}
println(p)
`
	tester.GeneralTestX(false, t, code, "{0 0}\n")
}
