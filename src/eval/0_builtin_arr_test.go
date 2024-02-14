package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestBuiltinArrayJoin(t *testing.T) {
	code := `
println([1, 2, 3].join(':'))
`
	expected := `1:2:3
`
	tester.GeneralTest(false, t, code, expected)
}

func TestBuiltinArrayPop(t *testing.T) {
	code := `
arr = [1, 2, 3]
arr, poped = arr.pop()
println(arr)
println(poped)
`
	expected := `[2 3]
1
`
	tester.GeneralTest(false, t, code, expected)
}
