package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
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
	tester.GeneralTest(false, t, code, "Ptr {x: 0, y: 0}\n")
}

func TestZeroValueUsable2(t *testing.T) {
	code := `
int p
println(p)
`
	tester.GeneralTest(true, t, code, "0\n")
}

func TestZeroValueUsable3(t *testing.T) {
	code := `
struct Ptr {
  int x
  int y
}

Ptr p
println(p)
`
	tester.GeneralTest(false, t, code, "Ptr {x: 0, y: 0}\n")
}
