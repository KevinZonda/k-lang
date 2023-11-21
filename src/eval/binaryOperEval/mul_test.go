package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestMul(t *testing.T) {
	tester.Assert(t, binaryOperEval.Mul(1, 2), 2)
	tester.Assert(t, binaryOperEval.Mul(1.1, 2), 2.2)
	tester.Assert(t, binaryOperEval.Mul(1, 2.1), 2.1)
}
