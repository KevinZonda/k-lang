package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestIfTrue(ts *testing.T) {
	code := `
if (1 == 1) {
	print("1")
} else {
	print("2")
}
`
	expected := `1`
	tester.GeneralTest(true, ts, code, expected)
}

func TestIfFalse(ts *testing.T) {
	code := `
if (1 == -2) {
	print("1")
} else {
	print("2")
}
`
	expected := `2`
	tester.GeneralTest(true, ts, code, expected)
}

func TestElseIf(ts *testing.T) {
	code := `
i = 3
if (i == 1) {
	println("1")
} else if (i == 2) {
	println("2")
} else {
	println("otherwise")
}
`
	expected := `otherwise
`
	tester.GeneralTest(true, ts, code, expected)
}
