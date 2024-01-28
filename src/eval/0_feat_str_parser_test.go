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

func TestStringParser2(t *testing.T) {
	code := `
z := 18
x := $"Hello World\n\nHi {z}"
print(x)
`
	expected := "Hello World\n\nHi 18"
	tester.GeneralTest(false, t, code, expected)
}

func TestStringParser3(t *testing.T) {
	code := `
z := 18
x := '\'\\ ""'
print(x)
`
	expected := "'\\ \"\""
	tester.GeneralTest(false, t, code, expected)
}
