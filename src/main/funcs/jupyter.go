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
	"strings"
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
		Banner:                fmt.Sprintf("Mini Language kernel - v%s", "0.0.1"),
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
	e       *eval.Eval
	context *eval.Eval
}

func (i *replBasedInterpreter) CompleteWords(code string, cursorPos int) (prefix string, completions []string, tail string) {
	return "", nil, ""
}

func (i *replBasedInterpreter) Eval(code string) (values []any, err error) {
	parser := parserHelper.FromString(code)

	if len(parser.Errors()) > 0 {
		sb := strings.Builder{}
		for _, e := range parser.Errors() {
			sb.WriteString("[Error] ")
			sb.WriteString(e.Error())
			sb.WriteString("\n")
		}
		return nil, fmt.Errorf(strings.TrimSpace(sb.String()))
	}

	ast := parser.Ast()

	e := eval.New(ast, "")
	e.LoadContext(i.context)
	if val, hasV := e.DoRetLastExpr(); hasV && val != nil {
		values = append(values, val)
	}
	i.context = e
	return values, nil
}

func newReplInterpreter() *replBasedInterpreter {
	return &replBasedInterpreter{
		e: eval.New(nil, ""),
	}
}

var _ jupyter.Interpreter = (*replBasedInterpreter)(nil)