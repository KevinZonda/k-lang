package minimalfunc

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli/consoleReader"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/jout"
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
		os.Exit(0)
	case ":clear", ":c":
		fmt.Print("\033[H\033[2J")
		fmt.Println(buildconst.Msg())
	case ":history":
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
	if input != "" {
		str, e := iox.ReadAllText(input)
		panicx.PanicIfNotNil(e, e)

		ast, _ := parserHelper.Ast(str)
		// TODO;
		r.context = eval.New(input)
		rst := r.context.DoSafely(ast)
		if rst.PrintPanic().IsPanic {
			os.Exit(1)
		}
	}
	fmt.Println(buildconst.Msg())

	rl, err := consoleReader.New("> ")
	panicx.PanicIfNotNil(err, err)

	for {
		line, err := rl.Readline()
		if err != nil {
			os.Exit(1)
		}

		if r.internalCmd(line) {
			continue
		}
		buffer, err := consoleReader.MultipleLine(rl, line, "> ", "| ")

		parser := parserHelper.FromString(buffer)

		if len(parser.Errors()) > 0 {
			for _, e := range parser.Errors() {
				fmt.Fprintln(os.Stderr, "[Error]", e.Error())
			}
			continue
		}

		r.history = append(r.history, buffer)
		ast := parser.Ast()
		if r.debug {
			for idx, node := range ast {
				fmt.Print("[", idx, "] -> ")
				jout.Println(node)
			}
		}
		if r.context == nil {
			r.context = eval.New("")
		}
		rst := r.context.DoSafely(ast)
		if rst.IsPanic {
			rst.PrintPanic()
		}

		if !r.dontPrintIt {
			if rst.HasReturn {
				fmt.Println("<", rst.ReturnValue)
			} else if rst.IsLastExpr {
				fmt.Println("<", rst.LastExprVal)
			}
		}

	}

}
