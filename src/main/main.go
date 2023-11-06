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
	// [from] [subCmd] [args...]
	args := parseCmdArgs()
	switch os.Args[1] {
	case "ast":
		funcs.Ast(args.GetStrOr(0, cli.Input), args.GetStrOr(1, cli.Output))
	case "repl":
		(&funcs.Repl{}).Repl()
	case "compile":
		funcs.Compile(args.GetStrOr(0, cli.Input), args.GetStrOr(1, cli.Output))
	case "run":
		funcs.Run(args.GetStrOr(0, cli.Input))
	default:
		if len(args) == 0 {
			panic("Not recognised command: " + os.Args[1])
		}
		if len(args) == 1 {
			funcs.Run(args[0])
		}
		panic("Not recognised command: " + strings.Join(os.Args[1:], " "))
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
	if *ss == "" {
		return ""
	}
	return *ss
}
