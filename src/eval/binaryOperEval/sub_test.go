package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestSub(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Sub, 1, 2), -1)
	tester.Assert(t, binaryOperEval.Sub(1.1, 2), float64(1.1)-float64(2))
	tester.Assert(t, binaryOperEval.BinaryOper(token.Sub, 1, 2.1), -1.1)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Sub, 1.1, 2.1), -1.0)
}
