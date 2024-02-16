package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestSetToConst(t *testing.T) {
	code := `
struct color {
    int r = 19
}
c := color{}
print(c.r)
`
	tester.GeneralTest(false, t, code, "19")
}

func TestSetToConst2(t *testing.T) {
	code := `
struct color {
    int r = 19
}
color = color{}
print(color.r)
`

	tester.ExpectPanic(t, code, func(exp string) bool {
		return exp == "cannot assign to color"
	})
}
