package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestOr(t *testing.T) {
	tester.Assert(t, binaryOperEval.Or(true, true), true)
	tester.Assert(t, binaryOperEval.Or(true, false), true)
	tester.Assert(t, binaryOperEval.Or(false, true), true)
	tester.Assert(t, binaryOperEval.Or(false, false), false)
}
