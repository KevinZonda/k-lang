package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestGreaterEq(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 1, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 1.0, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 1, 2), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 1.1, 1.1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 1.1, 1.2), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 1.5, 1.2), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 1.5, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 2, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.GreaterEq, 2, 1.1), true)

}

func TestGreater(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 1, 1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 1.0, 1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 1, 2), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 1.1, 1.1), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 1.1, 1.2), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 1.5, 1.2), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 1.5, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 2, 1), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Greater, 2, 1.1), true)
}
