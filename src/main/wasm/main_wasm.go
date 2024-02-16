package main

import (
	"bytes"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
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
	js.Global().Set("fmtCodeX", js.FuncOf(FmtCode))
	js.Global().Set("infoX", js.FuncOf(InfoX))
	<-c
}

func InfoX(this js.Value, args []js.Value) any {
	return js.ValueOf(buildconst.Msg())
}

func RunCodeX(this js.Value, args []js.Value) any {
	s := args[0].String()
	result := executeCode(s)
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

func executeCode(code string) string {
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		return parseErrors(errs)
	}
	e := eval.New(ast, "")
	stdout := &bytes.Buffer{}
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
