package main

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli/funcs"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/cli/shared"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/minimalfunc"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		(&minimalfunc.Repl{}).Repl("")
		return
	}
	// [from] [subCmd] [args...]
	args := parseCmdArgs()
	switch os.Args[1] {
	case "ast":
		funcs.Ast(args.GetStrOr(0, shared.Input), args.GetStrOr(1, shared.Output))
	case "repl":
		(&minimalfunc.Repl{}).Repl(args.GetStrOr(0, shared.Input))
	case "compile":
		funcs.Compile(args.GetStrOr(0, shared.Input), args.GetStrOr(1, shared.Output))
	case "run":
		minimalfunc.Run(args.GetStrOr(0, shared.Input))
	case "2ast":
		funcs.ShowAst(args.GetStrOr(0, shared.Input))
	case "help":
		funcs.Help()
	case "lsp":
		funcs.Lsp(args.GetStrOr(0, nil))
	case "fmt":
		funcs.Format(args.GetStrOr(0, shared.Input))
	case "jupyter":
		funcs.JupyterKernel(args.GetStrOr(0, nil))
	case "version", "info":
		fmt.Println(buildconst.Msg())
	case "mod":
		minimalfunc.Mod(args.GetStrOr(0, nil), args[1:])
	default:
		if len(args) == 0 {
			minimalfunc.Run(os.Args[1])
			return
		}
		fmt.Fprintln(os.Stderr, "Not recognised command: "+strings.Join(os.Args[1:], " "))
		fmt.Fprintln(os.Stderr, "Use `help` to see all commands")
		os.Exit(1)

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
