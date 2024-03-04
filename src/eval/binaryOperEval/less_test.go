package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestLessEq(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 1, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 1.0, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 1, 2), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 1.1, 1.1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 1.1, 1.2), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 1.5, 1.2), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 1.5, 1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 2, 1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.LessEq, 2, 1.1), false)

}

func TestLess(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 1, 1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 1.0, 1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 1, 2), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 1.1, 1.1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 1.1, 1.2), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 1.0, 1.2), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 0.7, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, 0, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Less, -1, 1.1), true)
}
