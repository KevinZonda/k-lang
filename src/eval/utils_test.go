package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/tx"
	"testing"
)

func generalTest(ts *testing.T, code, expected string) {
	if err := tx.IsStdoutAsExpected(func() {
		ast, _ := parserHelper.Ast(code)
		e := eval.New(ast)
		e.Do()
	}, expected); err != nil {
		ts.Fatal(err)
	}
}
