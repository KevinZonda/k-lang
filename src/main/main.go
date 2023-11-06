package main

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/funcs"
	"os"
	"strings"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli"
)

func main() {
	if len(os.Args) == 1 {
		(&funcs.Repl{}).Repl()
		return
	}
	switch os.Args[1] {
	case "ast":
		funcs.Ast(cli.GetStr(cli.Input), cli.GetStr(cli.Output))
	case "repl":
		(&funcs.Repl{}).Repl()
	case "compile":
		funcs.Compile(cli.GetStr(cli.Input), cli.GetStr(cli.Output))
	case "run":
		funcs.Run(cli.GetStr(cli.Input))
	default:
		panic("Not recognised command: " + strings.Join(os.Args[1:], " "))
	}
}
