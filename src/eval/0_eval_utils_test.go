package eval_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestEvalPanic(t *testing.T) {
	code := `
x : int = "114"
MEM()
continue
break
`
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		t.Fatal(errs)
		return
	}
	e := eval.New()
	fmt.Println(e.PtrAddr())
	e.LoadStdFromOS()
	tester.Assert(t, e.GetStdIn(), e.GetIO().GetStdin())
	tester.Assert(t, e.GetStdErr(), e.GetIO().GetStderr())
	tester.Assert(t, e.GetStdOut(), e.GetIO().GetStdout())

	e.GetMemory()
	tester.Assert(t, e.GetStdIn(), e.GetIO().GetStdin())
	tester.Assert(t, e.GetStdErr(), e.GetIO().GetStderr())
	tester.Assert(t, e.GetStdOut(), e.GetIO().GetStdout())

	e.SetStdIn(nil)
	e.SetStdErr(nil)
	e.SetStdOut(nil)
	e.ResetStd()
	tester.Assert(t, e.GetStdIn(), e.GetIO().GetStdin())
	tester.Assert(t, e.GetStdErr(), e.GetIO().GetStderr())
	tester.Assert(t, e.GetStdOut(), e.GetIO().GetStdout())

	e.WithVisualize()
	e.SetPath("")

	fmt.Println(e.String())
	rst := e.DoSafely(ast)

	fmt.Println(e.GetMemory().Bottom().String())
	fmt.Println(e.GetMemory().String())

	if rst.IsPanic {
		rst.PrintPanic()
		return
	}

}
