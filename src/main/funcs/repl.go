package funcs

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/consoleReader"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/jout"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"os"
	"reflect"
)

type Repl struct {
	debug   bool
	printIt bool
	context *eval.Eval
}

func (r *Repl) Repl() {
	rl, err := consoleReader.New("> ")
	panicx.PanicIfNotNil(err, err)

	var it any
	var history []string

	//defer func() {
	//	_ = rl.Close()
	//	if r := recover(); r != nil {
	//		fmt.Println("Recover from panic:", r)
	//		repl(context)
	//	}
	//}()
	for {
		line, err := rl.Readline()
		if err != nil {
			os.Exit(1)
		}

		switch line {
		case ":exit", ":quit", ":q", ":e":
			return
		case ":history":
			for i, h := range history {
				fmt.Println("[", i, "] ->", h)
			}
			continue
		case ":clear", ":c":
			r.context = nil
			continue
		case "+debug", "+d":
			r.debug = true
			fmt.Println("Debug: ON")
			continue
		case "-debug", "-d":
			r.debug = false
			fmt.Println("Debug: OFF")
			continue
		case "-it":
			r.printIt = false
			fmt.Println("Print it: OFF")
			continue
		case "+it":
			r.printIt = true
			fmt.Println("Print it: ON")
			continue
		case ":type", ":t":
			fmt.Println("Type: ", reflect.TypeOf(it))
			continue
		}

		buffer, err := consoleReader.MultipleLine(rl, line, "> ", "| ")

		parser := parserHelper.FromString(buffer)

		if len(parser.Errors()) > 0 {
			for _, e := range parser.Errors() {
				fmt.Println("[Error]", e.Error())
			}
			continue
		}

		history = append(history, buffer)
		ast := parser.Ast()
		if r.debug {
			for idx, node := range ast {
				fmt.Print("[", idx, "] -> ")
				jout.Println(node)
			}
		}
		e := eval.New(ast)
		e.LoadContext(r.context)
		e.Do()
		if r.printIt {
			fmt.Println("Evaluated ->", e.It())
		}
		it = e.It()
		r.context = e
	}

}
