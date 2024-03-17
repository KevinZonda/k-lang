package main

import (
	"bytes"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"io"
	"syscall/js"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/fmtr"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("runCodeX", js.FuncOf(RunCodeX))
	js.Global().Set("runCodeStreamX", js.FuncOf(RunCodeStream))
	js.Global().Set("fmtCodeX", js.FuncOf(FmtCode))
	js.Global().Set("infoX", js.FuncOf(InfoX))
	js.Global().Set("replX", js.FuncOf(RunReplX))
	js.Global().Set("initReplX", js.FuncOf(InitRepl))
	<-c
}

func InfoX(this js.Value, args []js.Value) any {
	return js.ValueOf(buildconst.Msg())
}

func RunCodeX(this js.Value, args []js.Value) any {
	s := args[0].String()
	stdout := &bytes.Buffer{}
	result := executeCode(stdout, s)
	return js.ValueOf(result)
}

func RunReplX(this js.Value, args []js.Value) any {
	s := args[0].String()
	stdout := &bytes.Buffer{}
	result := runRepl(stdout, s)
	return js.ValueOf([]any{stdout.String(), result})
}

func InitRepl(this js.Value, args []js.Value) any {
	initRepl()
	return js.Undefined()
}
func FmtCode(this js.Value, args []js.Value) any {
	s := args[0].String()
	result := fmtCode(s)
	return js.ValueOf(result)
}

func fmtCode(code string) string {
	codeX, err := fmtr.Fmt(code)
	if len(err) > 0 {
		return code
	}
	return codeX
}

type writer interface {
	io.Writer
	String() string
}

func executeCode(stdout writer, code string) string {
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		str := errs.String()
		stdout.Write([]byte(str))
		return str
	}
	e := eval.New()
	e.SetStdOut(stdout)
	e.SetStdErr(stdout)
	rst := e.DoSafely(ast)
	if rst.IsPanic {
		if stdout != nil {
			stdout.Write([]byte(rst.PanicString()))
		}
	}
	return stdout.String()
}

var repl *eval.Eval

func initRepl() {
	repl = eval.New()
}

func runRepl(stdout writer, code string) any {
	if repl == nil {
		initRepl()
	}
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		str := errs.String()
		stdout.Write([]byte(str))
		return str
	}
	e := repl
	e.SetStdOut(stdout)
	e.SetStdErr(stdout)
	rst := e.DoSafely(ast)
	if rst.IsPanic {
		if stdout != nil {
			stdout.Write([]byte(rst.PanicString()))
		}
		return nil
	}
	if rst.HasReturn {
		return rst.ReturnValue
	}
	return rst.LastExprVal
}

func RunCodeStream(this js.Value, args []js.Value) any {
	code := args[0].String()
	stdio := &jsWriter{
		setText: args[1],
		buf:     bytes.Buffer{},
	}

	executeCode(stdio, code)
	return js.Undefined()
}

type jsWriter struct {
	setText js.Value
	buf     bytes.Buffer
}

func (j *jsWriter) Write(p []byte) (n int, err error) {
	n, err = j.buf.Write(p)
	j.setText.Invoke(js.ValueOf(j.buf.String()))
	return len(p), nil
}

func (j *jsWriter) String() string {
	return j.buf.String()
}
