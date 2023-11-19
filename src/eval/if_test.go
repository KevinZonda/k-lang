package eval_test

import (
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
	GeneralTest(true, ts, code, expected)
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
	GeneralTest(true, ts, code, expected)
}
