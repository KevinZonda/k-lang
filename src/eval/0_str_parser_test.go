package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestStringParser(t *testing.T) {
	code := `
x:= @"Hello World\n\nHi {x}"
print(x)
`
	expected := "Hello World\\n\\nHi {x}"
	tester.GeneralTest(false, t, code, expected)
}

func TestStringParser1(t *testing.T) {
	code := `
x:= "Hello World\n\nHi {x}"
print(x)
`
	expected := "Hello World\n\nHi {x}"
	tester.GeneralTest(false, t, code, expected)
}
