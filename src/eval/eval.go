package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"io"
	"os"
	"path/filepath"
)

type Eval struct {
	ast          tree.Ast
	objTable     *TableStack
	basePath     string
	loopLvl      int
	currentToken token.Token
	builtin      builtin.BuiltIn
}

// region STD {IN, OUT, ERR}
func (e *Eval) SetStdOut(out io.WriteCloser) {
	e.builtin.StdOut = out
}

func (e *Eval) SetStdIn(in io.ReadCloser) {
	e.builtin.StdIn = in
}

func (e *Eval) SetStdErr(err io.WriteCloser) {
	e.builtin.StdErr = err
}

func (e *Eval) GetStdOut() io.WriteCloser {
	return e.builtin.StdOut
}

func (e *Eval) GetStdIn() io.ReadCloser {
	return e.builtin.StdIn
}

func (e *Eval) GetStdErr() io.WriteCloser {
	return e.builtin.StdErr
}

func (e *Eval) ResetStd() {
	e.builtin.StdOut = os.Stdout
	e.builtin.StdIn = os.Stdin
	e.builtin.StdErr = os.Stderr
}

//endregion

func (e *Eval) CurrentToken() token.Token {
	return e.currentToken
}

var openedFiles map[string]*Eval

func ResetGlobal() {
	openedFiles = make(map[string]*Eval)
}

func New(ast tree.Ast, inputFile string) *Eval {
	if openedFiles == nil {
		openedFiles = map[string]*Eval{}
	}
	path := filepath.Dir(inputFile)
	return &Eval{
		ast:      ast,
		objTable: NewObjectTable(),
		basePath: path,
		builtin:  builtin.NewBuiltIn(),
	}
}

func (e *Eval) LoadContext(o *Eval) {
	if o == nil {
		e.objTable = NewObjectTable()
		return
	}
	e.objTable = o.objTable
	if e.objTable.Empty() {
		e.objTable = NewObjectTable()
	}
	e.builtin = o.builtin
}

type runResult struct {
	retV           *obj.Object
	hasRet         bool
	lastExpr       any
	isLastElemExpr bool
}

func (r runResult) ReturnValue() any {
	if r.hasRet {
		return r.retV.Val
	}
	return nil
}

func (e *Eval) runAst(ast tree.Ast, breaks ...string) runResult {
	result := runResult{}
	for idx, n := range ast {
		for _, key := range breaks {
			if e.objTable.HasKeyAtTop(key) {
				goto end
			}
		}
		switch nT := n.(type) {
		case *node.CodeBlock:
			e.EvalCodeBlock(nT)
		case node.Expr:
			exprR := e.EvalExpr(nT)
			if idx == len(e.ast)-1 {
				switch n.(type) {
				case *node.AssignStmt:
					result.isLastElemExpr = false
				default:
					result.isLastElemExpr = true
					result.lastExpr = exprR
				}
			}
		case node.Stmt:
			e.EvalStmt(nT)
		case *node.FuncBlock:
			e.objTable.Set(nT.Name.Value, nT)
		case *node.StructBlock:
			e.objTable.Set(nT.Name, nT)
		case *node.OpenBlock:
			for _, stmt := range nT.Openers {
				e.EvalStmt(stmt)
			}
		default:
			panic("not implemented")
		}
	}
end:
	retV, hasRet := e.objTable.Get(reserved.Return)
	result.retV = retV
	result.hasRet = hasRet
	return result
}

func (e *Eval) runWithBreak(breaks ...string) runResult {
	return e.runAst(e.ast, breaks...)
}

func (e *Eval) Do() (retV any, hasRet bool) {
	if r := e.runWithBreak(reserved.Return); r.hasRet {
		return r.retV, r.hasRet
	}

	retV = e.EvalMain()
	hasRet = retV != nil
	return
}

func (e *Eval) DoSafely() (rst DetailedRunResult) {
	defer func() {
		if r := recover(); r != nil {
			rst.IsPanic = true
			rst.PanicMsg = fmt.Sprint(r)
			rst.CurrentToken = e.CurrentToken()
		}
	}()

	rst.stderr = e.GetStdErr()
	rst.ReturnValue, rst.HasReturn = e.Do()
	rst.CurrentToken = e.CurrentToken()
	return
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

func (e *Eval) DoRetLastExpr() (lastExprVal any, isLastExpr bool) {
	rr := e.runWithBreak(reserved.Return)
	return rr.lastExpr, rr.isLastElemExpr
}

func (e *Eval) EvalMain() any {
	fn, ok := e.objTable.Get("main")
	if !ok || !fn.Is(obj.Func) {
		return nil
	}
	fx := fn.ToFunc()
	e.frameStart(true)

	result := e.runAst((tree.Ast)(fx.Body.Nodes), reserved.Return)

	e.frameEnd()

	return result.ReturnValue()
}

func (e *Eval) PtrAddr() string {
	return fmt.Sprintf("%p", e)
}

func (e *Eval) String() string {
	return fmt.Sprintf("Eval@%p", e)
}
