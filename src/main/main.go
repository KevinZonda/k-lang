package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/consoleReader"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/compressor"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
)

var debug bool = false
var printIt bool = false

func main() {
	cli.ParseParam()
	switch cli.Ctx().Command() {
	case "ast <input> <output>":
		ast(cli.Param().Ast.Input, cli.Param().Ast.Output)
	case "repl":
		repl(nil)
	case "compile <input> <output>":
		compile(cli.Param().Compile.Input, cli.Param().Compile.Output)
	case "run <input>":
		run(cli.Param().Run.Input)
	default:
		panic(cli.Ctx().Command())
	}
}

func compile(input string, output string) {
	if output == "" {
		output = strings.TrimSuffix(input, filepath.Ext(input)) + ".ast"
	}

	str, e := iox.ReadAllText(input)
	if e != nil {
		panic(e)
	}
	ast := parserHelper.Ast(str)
	e = iox.WriteAllBytes(output, compressor.Compress(ast))
	if e != nil {
		panic(e)
	}
}

func run(input string) {
	bs, e := iox.ReadAllByte(input)
	if e != nil {
		panic(e)
	}
	ast := compressor.Decompress(bs)
	ev := eval.New(ast)
	ev.Do()
	fmt.Println("Evaluated ->", ev.It())
}

func repl(context *eval.Eval) {
	rl, err := consoleReader.New("> ")
	var it any
	var history []string
	if err != nil {
		panic(err)
	}
	defer func() {
		rl.Close()
		if r := recover(); r != nil {
			fmt.Println("Recover from panic:", r)
			repl(context)
		}
	}()
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
			context = nil
			continue
		case "+debug", "+d":
			debug = true
			fmt.Println("Debug: ON")
			continue
		case "-debug", "-d":
			debug = false
			fmt.Println("Debug: OFF")
			continue
		case "-it":
			printIt = false
			fmt.Println("Print it: OFF")
			continue
		case "+it":
			printIt = true
			fmt.Println("Print it: ON")
			continue
		case ":type", ":t":
			fmt.Println("Type: ", reflect.TypeOf(it))
			continue
		}

		buffer := ""

		if strings.HasSuffix(line, "\\") {
			for strings.HasSuffix(line, "\\") {
				buffer += strings.TrimSuffix(line, "\\") + "\n"
				rl.SetPrompt("| ")
				line, err = rl.Readline()
				if err != nil {
					panic(err)
				}
			}
			buffer += line
			rl.SetPrompt("> ")
		} else {
			buffer = line
		}

		parser := parserHelper.FromString(buffer)

		if len(parser.Errors()) > 0 {
			for _, e := range parser.Errors() {
				fmt.Println("Error:", e.Error())
			}
			continue
		}

		history = append(history, buffer)
		ast := parser.Ast()
		if debug {
			for idx, node := range ast {
				bs, _ := json.MarshalIndent(node, "", "    ")
				fmt.Println(idx, "->", string(bs))
			}
		}
		e := eval.New(ast)
		e.LoadContext(context)
		e.Do()
		if printIt {
			fmt.Println("Evaluated ->", e.It())
		}
		it = e.It()
		context = e
	}

}

func ast(input string, output string) {
	txt, err := iox.ReadAllText(input)
	if err != nil {
		panic(err)
	}
	ast := parserHelper.Ast(txt)
	bs, err := json.MarshalIndent(ast, "", "    ")
	if err != nil {
		panic(err)
	}
	err = iox.WriteAllText(output, string(bs))
}
