package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Pow, 1, 2), 1)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Pow, 1.1, 2), math.Pow(1.1, 2))
	tester.Assert(t, binaryOperEval.BinaryOper(token.Pow, 1, 2.1), 1.0)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Pow, 2, 2), 4)
}
