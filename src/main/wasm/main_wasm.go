package main

import (
	"bytes"
	"strconv"
	"strings"
	"syscall/js"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("runCode", js.FuncOf(RunCode))
	js.Global().Set("runCodeX", js.FuncOf(RunCodeX))
	<-c
}

func RunCode(this js.Value, args []js.Value) any {
	return js.ValueOf("Call from RunCode")
}

func RunCodeX(this js.Value, args []js.Value) any {
	s := args[0].String()
	result := executeCode(s)
	return js.ValueOf(result)
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
		return rst.PanicMsg
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
