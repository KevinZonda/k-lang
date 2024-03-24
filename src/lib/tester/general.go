package tester

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"strings"
	"testing"
)

type Pair struct {
	Code       string
	Expected   string
	ExpectFunc func(string) bool
	AllowPanic bool
}

func BatchRun(t *testing.T, pairs ...Pair) {
	for _, pair := range pairs {
		if pair.ExpectFunc != nil {
			GeneralTestLambda(pair.AllowPanic, t, pair.Code, pair.ExpectFunc)
		} else {
			GeneralTest(pair.AllowPanic, t, pair.Code, pair.Expected)
		}
		if t.Failed() {
			return
		}
	}
}

func BatchRunSplit(t *testing.T, tests []string, expected []string) {
	if len(tests) != len(expected) {
		t.Fatal("len(tests) != len(expected)")
		return
	}
	for i := 0; i < len(tests); i++ {
		GeneralTest(false, t, tests[i], expected[i])
		if t.Failed() {
			return
		}
	}
}

func Strings(s ...string) []string {
	return s
}

func GeneralTest(allowPanic bool, ts *testing.T, code, expected string) {
	if err := IsStdoutAsExpected(allowPanic, func() {
		ast, errs := parserHelper.Ast(code)
		if len(errs) > 0 {
			ts.Fatal(errs)
			return
		}
		e := eval.New()
		e.Do(ast)
	}, expected); err != nil {
		ts.Fatal(err)
	}
}

func ExpectPanic(ts *testing.T, code string, expected func(exp string) bool) {
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		ts.Fatal(errs)
		return
	}
	e := eval.New()
	rst := e.DoSafely(ast)
	if !rst.IsPanic {
		ts.Fatal("NOT PANIC")
		return
	}
	ts.Log("PANIC WITH:", rst.PanicMsg)
	if !expected(rst.PanicMsg) {
		ts.Fatalf("got %s", rst.PanicMsg)
		return
	}

}

func PanicFx(ts *testing.T, code func(), expected func(exp string) bool) {
	defer func() {
		if r := recover(); r != nil {
			if !expected(r.(string)) {
				ts.Fatalf("got %s", r)
			}
		}
	}()
	code()
}

func PanicExFxNotEmptyString(exp string) bool {
	return exp != ""
}
func NoPanic(ts *testing.T, code string) {
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		ts.Fatal(errs)
		return
	}
	e := eval.New()
	rst := e.DoSafely(ast)
	if rst.IsPanic {
		ts.Fatal("PANIC WITH:", rst.PanicMsg)
		return
	}
}

func GeneralTestLambda(allowPanic bool, ts *testing.T, code string, expected func(string) bool) {
	op := Stdout(allowPanic, func() {
		ast, errs := parserHelper.Ast(code)
		if len(errs) > 0 {
			ts.Fatal(errs)
			return
		}
		e := eval.New()
		e.Do(ast)
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
	e := eval.New()
	e.Do(ast)
}

func ContainsAll(str string, subs ...string) bool {
	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			return false
		}
	}
	return true
}
