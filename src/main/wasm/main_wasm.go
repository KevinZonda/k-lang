package main

import (
	"bytes"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"io"
	"strconv"
	"strings"
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
		return parseErrors(errs)
	}
	e := eval.New(ast, "")
	e.SetStdOut(stdout)
	e.SetStdErr(stdout)
	rst := e.DoSafely()
	if rst.IsPanic {
		rst.PrintPanic()
	}
	return stdout.String()
}

func parseErrors(errs []parserHelper.CodeError) string {
	sb := strings.Builder{}
	for idx, err := range errs {
		sb.WriteString("[" + strconv.Itoa(idx) + "] ")
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}
	return sb.String()
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
	js.Global().Call("setResult", js.ValueOf(j.buf.String()))
	return len(p), nil
}

func (j *jsWriter) String() string {
	return j.buf.String()
}
