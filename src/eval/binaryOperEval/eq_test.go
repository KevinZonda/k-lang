package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestEq(t *testing.T) {
	tester.Assert(t, binaryOperEval.Eq(true, true), true)
	tester.Assert(t, binaryOperEval.Eq(true, false), false)
	tester.Assert(t, binaryOperEval.Eq(false, true), false)
	tester.Assert(t, binaryOperEval.Eq(false, false), true)
	tester.Assert(t, binaryOperEval.Eq(1, 1), true)
	tester.Assert(t, binaryOperEval.Eq(1.0, 1), true)
	tester.Assert(t, binaryOperEval.Eq(1, 2), false)
	tester.Assert(t, binaryOperEval.Eq(1.1, 1.1), true)
	tester.Assert(t, binaryOperEval.Eq(1.1, 1.2), false)
	tester.Assert(t, binaryOperEval.Eq("1", "1"), true)
	tester.Assert(t, binaryOperEval.Eq("1", "2"), false)
	tester.Assert(t, binaryOperEval.Eq("1.1", "1.1"), true)
	tester.Assert(t, binaryOperEval.Eq("1.1", "1.2"), false)
}
