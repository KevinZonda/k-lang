package minimalfunc

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/jout"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli/consoleReader"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"os"
)

type Repl struct {
	debug       bool
	dontPrintIt bool
	context     *eval.Eval
	history     []string
}

func (r *Repl) internalCmd(line string) bool {
	switch line {
	case ":exit", ":quit", ":q", ":e":
		buildconst.Exit(0)
	case ":clear", ":c":
		fmt.Print("\033[H\033[2J")
		fmt.Println(buildconst.Msg())
	case ":history", ":h":
		for idx, h := range r.history {
			fmt.Println("[", idx, "] -> ", h)
		}

	case "+debug", "+d":
		r.debug = true
		fmt.Println("Debug: ON")

	case "-debug", "-d":
		r.debug = false
		fmt.Println("Debug: OFF")
	case "-it":
		r.dontPrintIt = true
		fmt.Println("Print it: OFF")

	case "+it":
		r.dontPrintIt = false
		fmt.Println("Print it: ON")
	default:
		return false
	}
	return true
}

func (r *Repl) Repl(input string) {
	r.context = eval.New().WithBasePath(input)

	if input != "" {
		str, e := iox.ReadAllText(input)
		panicx.NotNilErr(e)

		ast, errs := parserHelper.Ast(str)
		if len(errs) > 0 {
			fmt.Fprintln(os.Stderr, errs.String())
		} else if r.context.DoSafely(ast).PrintPanic().IsPanic {
			buildconst.Exit(1)
		}
	}
	fmt.Println(buildconst.Msg())

	rl, err := consoleReader.New("> ")
	panicx.NotNilErr(err)

	for {
		line, err := rl.Readline()
		if err != nil {
			buildconst.Exit(1)
			return
		}

		if r.internalCmd(line) {
			continue
		}
		buffer, err := consoleReader.MultipleLine(rl, line, "> ", "| ")

		ast, errs := parserHelper.Ast(buffer)
		if len(errs) > 0 {
			fmt.Fprintln(os.Stderr, errs.String())
			continue
		}
		r.history = append(r.history, buffer)
		if r.debug {
			for idx, node := range ast {
				fmt.Print("[", idx, "] -> ")
				jout.Println(node)
			}
		}
		rst := r.context.DoSafely(ast)
		if rst.IsPanic {
			rst.PrintPanic()
		}

		if !r.dontPrintIt {
			if val, ok := rst.VizValue(); ok {
				fmt.Println("<", val)
			}
		}

	}

}
