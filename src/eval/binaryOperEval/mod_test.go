package binaryOperEval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/binaryOperEval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestMod(t *testing.T) {
	tester.Assert(t, binaryOperEval.BinaryOper(token.Mod, 12, 7), 5)
	//tester.Assert(t, binaryOperEval.BinaryOper(token.Mod, 12.1, 7), 5.1)
	//tester.Assert(t, binaryOperEval.BinaryOper(token.Mod, 12, 7.1), 4.9)
}
