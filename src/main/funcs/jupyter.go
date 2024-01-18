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
		Banner:                fmt.Sprintf("Mini Language kernel - v%s", "0.0.1"),
		LanguageInfo: jupyter.KernelLanguageInfo{
			Name:           "K",
			Version:        runtime.Version(),
			FileExtension:  ".k",
			CodeMirrorMode: "application/x-cypher-query",
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
		for _, e := range parser.Errors() {
			fmt.Println("[Error]", e.Error())
		}
		return nil, nil
	}

	ast := parser.Ast()

	e := eval.New(ast, "")
	e.LoadContext(i.context)
	e.Do()
	i.context = e
	return nil, nil
}

func newReplInterpreter() *replBasedInterpreter {
	return &replBasedInterpreter{
		e: eval.New(nil, ""),
	}
}

var _ jupyter.Interpreter = (*replBasedInterpreter)(nil)
