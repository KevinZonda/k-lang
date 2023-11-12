package eval_test

import "testing"

func TestLambdaCall(t *testing.T) {
	code := `
foo := (x) {
	println(x+11)
}
foo(14)
`
	expected := "25\n"
	generalTest(t, code, expected)
}

func TestReturnLambda(t *testing.T) {
	code := `
foo := () {
	return (x) {
        println(x+11)
    }
}
foo()(14)
`
	expected := "25\n"
	generalTest(t, code, expected)
}
