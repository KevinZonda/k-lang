package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestEq(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, nil, nil), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, true, true), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, true, false), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, false, true), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, false, false), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, 1, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, 1.0, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, 1, 1.0), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, 1, 2), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, 1.1, 1.1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, 1.1, 1.2), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, "1", "1"), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, "1", "2"), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, "1.1", "1.1"), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, "1.1", "1.2"), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, []any{1, 2}, []any{1, 2}), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, []any{1, 2}, []any{1, 3}), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.NotEq, []any{1, 2}, []any{1, 3}), true)

	var a float64 = .1
	var b float64 = .2
	tester.Assert(t, binaryOperEval.BinaryOper(token.Equals, a+b, .3), true)
}
