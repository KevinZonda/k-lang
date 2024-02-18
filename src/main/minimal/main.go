package main

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/minimalfunc"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		(&minimalfunc.Repl{}).Repl("")
		return
	}
	switch args[0] {
	case "version", "info":
		fmt.Println(buildconst.Msg())
	case "run":
		minimalfunc.Run(args[1])
	default:
		minimalfunc.Run(args[0])
	}
}
