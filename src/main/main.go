package main

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/funcs"
	"os"
	"strings"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli"
)

func main() {
	if len(os.Args) == 1 {
		(&funcs.Repl{}).Repl("")
		return
	}
	// [from] [subCmd] [args...]
	args := parseCmdArgs()
	switch os.Args[1] {
	case "ast":
		funcs.Ast(args.GetStrOr(0, cli.Input), args.GetStrOr(1, cli.Output))
	case "repl":
		(&funcs.Repl{}).Repl(args.GetStrOr(0, cli.Input))
	case "compile":
		funcs.Compile(args.GetStrOr(0, cli.Input), args.GetStrOr(1, cli.Output))
	case "run":
		funcs.Run(args.GetStrOr(0, cli.Input))
	case "2ast":
		funcs.ShowAst(args.GetStrOr(0, cli.Input))
	case "help":
		funcs.Help()
	case "lsp":
		funcs.Lsp(args.GetStrOr(0, nil))
	default:
		if len(args) == 0 {
			funcs.Run(os.Args[1])
		} else {
			fmt.Println("Not recognised command: " + strings.Join(os.Args[1:], " "))
			fmt.Println("Use `help` to see all commands")
			os.Exit(1)
		}
	}
}

type cmdArgs []string

func parseCmdArgs() cmdArgs {
	if len(os.Args) <= 2 {
		return nil
	}
	return cmdArgs(os.Args[2:])
}

func (a cmdArgs) GetStrOr(i int, ss *string) string {
	if len(a) > i {
		if a[i] != "" {
			return a[i]
		}
	}
	if ss == nil {
		return ""
	}
	if *ss == "" {
		return ""
	}
	return *ss
}
