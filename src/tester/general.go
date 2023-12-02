package tester

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"strings"
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

func GeneralTestLambda(allowPanic bool, ts *testing.T, code string, expected func(string) bool) {
	op := Stdout(allowPanic, func() {
		ast, errs := parserHelper.Ast(code)
		if len(errs) > 0 {
			ts.Fatal(errs)
			return
		}
		e := eval.New(ast, "")
		e.Do()
	})
	if !expected(op) {
		ts.Fatalf("expected failed, got: %v", op)
	}
}

func GeneralTestX(allowPanic bool, ts *testing.T, code, expected string) {
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		ts.Fatal(errs)
		return
	}
	e := eval.New(ast, "")
	e.Do()
}

func ContainsAll(str string, subs ...string) bool {
	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			return false
		}
	}
	return true
}
