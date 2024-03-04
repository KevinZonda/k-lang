package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestDiv(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Div, 1, 2), 0.5)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Div, 1.1, 2), 0.55)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Div, 1, 2.2), 1/2.2)
	tester.Assert(t, binaryOperEval.BinaryOper(token.Div, 2.2, 1.1), 2.2/1.1)
}
