package tester

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func GeneralTest(allowPanic bool, ts *testing.T, code, expected string) {
	if err := IsStdoutAsExpected(allowPanic, func() {
		ast, errs := parserHelper.Ast(code)
		if len(errs) > 0 {
			ts.Fatal(errs)
			return
		}
		e := eval.New(ast, "")
		e.Do()
	}, expected); err != nil {
		ts.Fatal(err)
	}
}
