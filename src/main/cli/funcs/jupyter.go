package funcs

import (
	"encoding/json"
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/go-jupyter"
	"log"
	"runtime"
)

func JupyterKernel(file string) {
	var connInfo jupyter.ConnectionInfo
	{
		connData, err := iox.ReadAllByte(file)
		if err != nil {
			log.Fatal(err)
		}

		if err = json.Unmarshal(connData, &connInfo); err != nil {
			log.Fatal(err)
		}
	}

	jupyter.RunKernel(newReplInterpreter(), connInfo, jupyter.KernelInfo{
		ProtocolVersion:       jupyter.ProtocolVersion,
		Implementation:        "minikernel",
		ImplementationVersion: "0.0.1",
		Banner:                fmt.Sprintf("K Language kernel - v%s", "0.0.1"),
		LanguageInfo: jupyter.KernelLanguageInfo{
			Name:           "K",
			Version:        runtime.Version(),
			FileExtension:  ".k",
			CodeMirrorMode: "text/x-julia",
		},
		HelpLinks: []jupyter.KernelInfoHelpLink{},
	})
}

type replBasedInterpreter struct {
	e *eval.Eval
}

func (i *replBasedInterpreter) CompleteWords(code string, cursorPos int) (prefix string, completions []string, tail string) {
	return "", nil, ""
}

func (i *replBasedInterpreter) Eval(code string) (values []any, err error) {
	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		return nil, fmt.Errorf(errs.String())
	}

	if i.e == nil {
		i.e = eval.New().WithVisualize()
	}
	e := i.e

	e.LoadStdFromOS()

	rst := e.DoSafely(ast)
	if rst.IsPanic {
		return nil, fmt.Errorf(rst.PanicMsg)
	}
	val, ok := rst.VizValue()
	if ok && val != nil {
		values = append(values, val)
	}
	return values, nil
}

func newReplInterpreter() *replBasedInterpreter {
	return &replBasedInterpreter{
		e: eval.New().WithVisualize(),
	}
}

var _ jupyter.Interpreter = (*replBasedInterpreter)(nil)
