package builtin

import (
	"errors"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"os/exec"
)

type StdExecLib struct {
	FBLibrary
}

func NewStdExecLib() *StdExecLib {
	cmd := FxToFuncBlock(exec_cmd)
	cmd.Args = []*node.FuncArg{&node.FuncArg{Name: node.NewIdentifier(""), Type: NewTypeN("string"), Param: true}}
	return &StdExecLib{
		FBLibrary: FBLibrary{
			F: map[string]*node.FuncBlock{
				"cmd": cmd,
			},
		},
	}
}

func exec_cmd(args []any) string {
	var _args []string
	for _, arg := range args {
		_args = append(_args, arg.(string))
	}
	cmd := exec.Command(_args[0], _args[1:]...)
	out, err := cmd.Output()
	if err != nil {
		// FIXME: Panic?
		// Q
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			// log.Println(exitError.ExitCode())
		}
		panic(err)
	}
	return string(out)

}
