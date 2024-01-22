package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestStackframeProtection(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str, ok := r.(string)
			if !ok {
				t.Errorf("Expected string, got %T", r)
				return
			}
			t.Log(str)
			if str != "No Var Found: x" {
				t.Errorf("Expected 'No Var Found: x', got %s", str)
			}

		} else {
			t.Errorf("Expected panic, got %v", r)
		}
	}()
	code := `
fn f() {
	print(y)
    x = 18
    f2()
}
fn f2() {
  print(x)
}
y = 11
f()
`
	expected := ""
	tester.GeneralTest(true, t, code, expected)
}

func TestStackframeProtection2(t *testing.T) {
	code := `
fn f(i) {
	i = i + 1
	println(i)
}
i = 10
println(i)
f(i)
println(i)
`
	expected := "10\n11\n10\n"
	tester.GeneralTest(true, t, code, expected)
}
