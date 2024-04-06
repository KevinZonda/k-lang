package eval_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"os"
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
	tester.Assert(t, e.GetIO().GetStdIn(), os.Stdin)
	tester.Assert(t, e.GetIO().GetStdErr(), os.Stderr)
	tester.Assert(t, e.GetIO().GetStdOut(), os.Stdout)

	e.GetMemory()

	e.SetStdIn(nil)
	e.SetStdErr(nil)
	e.SetStdOut(nil)
	tester.Assert(t, e.GetIO().GetStdIn(), os.Stdin)
	tester.Assert(t, e.GetIO().GetStdErr(), os.Stderr)
	tester.Assert(t, e.GetIO().GetStdOut(), os.Stdout)
	e.LoadStdFromOS()
	tester.Assert(t, e.GetIO().GetStdIn(), os.Stdin)
	tester.Assert(t, e.GetIO().GetStdErr(), os.Stderr)
	tester.Assert(t, e.GetIO().GetStdOut(), os.Stdout)

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
