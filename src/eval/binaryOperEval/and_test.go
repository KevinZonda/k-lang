package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestAnd(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.And, true, true), true)
	tester.Assert(t, binaryOperEval.BinaryOper(token.And, true, false), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.And, false, true), false)
	tester.Assert(t, binaryOperEval.BinaryOper(token.And, false, false), false)
}
