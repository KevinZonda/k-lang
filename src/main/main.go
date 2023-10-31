package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/chzyer/readline"
)

var debug bool = false

func main() {
	cli.ParseParam()
	switch cli.Ctx().Command() {
	case "ast <input> <output>":
		ast(cli.Param().Ast.Input, cli.Param().Ast.Output)
	case "repl":
		repl("")
	default:
		panic(cli.Ctx().Command())
	}

}

func repl(context string) {
	rl, err := readline.New("> ")
	var it any
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
		case ":context", ":ctx":
			fmt.Println(context)
			continue
		case ":clear", ":c":
			context = ""
			continue
		case "+debug", "+d":
			debug = true
			fmt.Println("Debug: ON")
			continue
		case "-debug", "-d":
			debug = false
			fmt.Println("Debug: OFF")
			continue
		case ":type", ":t":
			fmt.Println("Type: ", reflect.TypeOf(it))
			continue
		}

		buffer := context + "\n" + line
		parser := parserHelper.FromString(buffer)
		ast := parser.Ast()
		if len(parser.Errors()) > 0 {
			for _, e := range parser.Errors() {
				fmt.Println("Error:", e.Error())
			}
			continue
		}
		if debug {
			for idx, node := range ast {
				bs, _ := json.MarshalIndent(node, "", "    ")
				fmt.Println(idx, "->", string(bs))
			}
		}
		e := eval.New(ast)
		e.Do()
		fmt.Println("Evaluated ->", e.It())
		it = e.It()
		context = buffer
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