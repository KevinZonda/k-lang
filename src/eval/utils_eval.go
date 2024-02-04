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
	IsLastExpr   bool
	stderr       io.Writer
}

func (rst DetailedRunResult) PrintPanic() DetailedRunResult {
	if !rst.IsPanic {
		return rst
	}
	tk := rst.CurrentToken
	stde := rst.stderr
	if stde == nil {
		stde = os.Stderr
	}
	fmt.Fprintf(rst.stderr, "%s at position L%d,%d-L%d,%d\n", rst.PanicMsg, tk.BeginLine, tk.BeginColumn, tk.EndLine, tk.EndColumn)
	return rst
}
