package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestMul(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Mul, 1, 2), 2)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Mul, 1.1, 2), 2.2)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Mul, 1, 2.1), 2.1)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Mul, 1.1, 2.1), float64(1.1)*float64(2.1))

}
