package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestEvalStructLiteral(t *testing.T) {
	src := `
x := colour {
   red: 10,
   green: 20,
   blue: 30,
}
println(x)
	`
	tester.GeneralTestX(true, t, src, "colour{red:10, green:20, blue:30}")
}
