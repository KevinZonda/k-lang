package main

import (
	"encoding/json"
	"fmt"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/eval"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/parserHelper"
	"github.com/chzyer/readline"
)

var debug bool = false

func main() {
	repl()
}

func repl() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	buffer := ""
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		switch line {
		case "exit":
			return
		case "clear":
			buffer = ""
			continue
		case "+debug":
			debug = true
			fmt.Println("Debug: ON")
			continue
		case "-debug":
			debug = false
			fmt.Println("Debug: OFF")
			continue
		}

		buffer += "\n" + line
		ast := parserHelper.Ast(buffer)
		if debug {
			for idx, node := range ast {
				bs, _ := json.MarshalIndent(node, "", "    ")
				fmt.Println(idx, "->", string(bs))
			}
		}
		e := eval.New(ast)
		e.Do()
		fmt.Println("Evaluated ->", e.It())
	}

}
