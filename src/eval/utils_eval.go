package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"io"
	"os"
)

func (e *Eval) PtrAddr() string {
	return fmt.Sprintf("%p", e)
}

func (e *Eval) String() string {
	return fmt.Sprintf("Eval@%p", e)
}

type DetailedRunResult struct {
	ReturnValue  any
	ReturnObj    *obj.Object
	HasReturn    bool
	IsPanic      bool
	PanicMsg     string
	CurrentToken token.Token
	LastExprVal  any
	HasLastExpr  bool
	stderr       io.Writer
}

// value returns the value of the last expression or the return value of the function
func (rst DetailedRunResult) value() (val any, ok bool) {
	if rst.HasReturn {
		return rst.ReturnValue, true
	}
	if rst.HasLastExpr {
		return rst.LastExprVal, true
	}
	return nil, false
}

func (rst DetailedRunResult) VizValue() (val any, ok bool) {
	valA, ok := rst.value()
	if !ok {
		return "", false
	}
	if obj.VizAny {
		return obj.TreeAnyT("", valA, false), true
	}
	return fmt.Sprint(valA), true
}

func (rst DetailedRunResult) PrintPanic() DetailedRunResult {
	if !rst.IsPanic {
		return rst
	}
	stde := rst.stderr
	if stde == nil {
		stde = os.Stderr
	}
	fmt.Fprintln(stde, rst.PanicString())
	return rst
}

func (rst DetailedRunResult) PanicString() string {
	if !rst.IsPanic {
		return ""
	}
	tk := rst.CurrentToken
	return fmt.Sprintf("%s at position L%d,%d-L%d,%d\n", rst.PanicMsg, tk.BeginLine, tk.BeginColumn, tk.EndLine, tk.EndColumn)
}
