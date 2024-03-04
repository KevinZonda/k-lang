package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestAdd(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, 1, 2), 3)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, 1.1, 2), 3.1)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, 1, 2.1), 3.1)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, 1.1, 2), 3.1)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, 1.1, 2.1), 3.2)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, "1", "2"), "12")
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, "1", 2), "12")
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, "1.1", "2"), "1.12")
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, "1", "2.1"), "12.1")
	tester.Assert(t, binaryOperEval.BinaryOper(token.Add, 1, "2.1"), "12.1")
	tester.AssertArray(t, binaryOperEval.BinaryOper(token.Add, []any{1}, []any{2}).([]any), []any{1, 2})
	tester.AssertArray(t, binaryOperEval.BinaryOper(token.Add, []any{1}, 2).([]any), []any{1, 2})
	// TODO: Fix this
	//tester.Assert(t, binaryOperEval.Add(1, "2.1"), "12.1")
	//tester.Assert(t, binaryOperEval.Add("1.1", 2), "1.12")
	//tester.Assert(t, binaryOperEval.Add(1.1, "2"), "1.12")
}
