package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestFuncCallPassVar(t *testing.T) {
	code := `
fn tailFact(n : int) int {
    println(n-1, n)
   return factT(n-1, n)
}

fn factT(n : int, current : int) int {
    println(n, current)
    return 12;
}

println(tailFact(20))
`
	tester.GeneralTest(false, t, code, "1920\n1920\n12\n")
}
