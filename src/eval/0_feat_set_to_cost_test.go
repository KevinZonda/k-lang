package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestSetToConst(t *testing.T) {
	code := `
struct color {
    r int = 19
}
color := color{}
print(color.r)
`
	tester.GeneralTestX(false, t, code, "19")
}

func TestSetToConst2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			x := r.(string)
			if x != "cannot assign to color" {
				t.Fail()
			}
		}
	}()
	code := `
struct color {
    r int = 19
}
color = color{}
print(color.r)
`

	tester.GeneralTestX(true, t, code, "19")
}
