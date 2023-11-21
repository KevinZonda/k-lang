package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestDiv(t *testing.T) {
	tester.Assert(t, binaryOperEval.Div(1, 2), 0.5)
	tester.Assert(t, binaryOperEval.Div(1.1, 2), 0.55)
}
