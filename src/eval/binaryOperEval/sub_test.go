package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestSub(t *testing.T) {
	tester.Assert(t, binaryOperEval.Sub(1, 2), -1.0)
	tester.Assert(t, binaryOperEval.Sub(1.1, 2), -0.9)
	tester.Assert(t, binaryOperEval.Sub(1, 2.1), -1.1)
	tester.Assert(t, binaryOperEval.Sub(1.1, 2.1), -1.0)
}
