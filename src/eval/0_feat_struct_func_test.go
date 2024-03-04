package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestFuncInStruct(t *testing.T) {
	code := `
struct x {
  int y
  fn x() {
    println("CALL FROM FUNC")
    self.y = 114
  }
}
y = x{}
y.x()
println(y.y)
`
	expected := `CALL FROM FUNC
114
`
	tester.GeneralTest(false, t, code, expected)
}
