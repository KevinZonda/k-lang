package eval_test

import (
	"testing"
)

func TestLambdaCall(t *testing.T) {
	code := `
foo := fn (x) {
	println(x+11)
}
foo(14)
`
	expected := "25\n"
	GeneralTest(false, t, code, expected)
}

func TestReturnLambda(t *testing.T) {
	code := `
foo := fn () {
	return fn (x) {
        println(x+11)
        return x + 11
    }
}
m := foo()
k := m(17)
`
	expected := "28\n"
	GeneralTest(false, t, code, expected)
}
