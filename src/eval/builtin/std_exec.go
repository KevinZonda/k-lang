package builtin

import (
	"os/exec"
)

type StdExecLib struct{}

func NewStdExecLib() *StdExecLib {
	return &StdExecLib{}
}

func (s *StdExecLib) FuncCall(caller string, args []any) any {
	switch caller {
	case "cmd":
		var _args []string
		for _, arg := range args {
			_args = append(_args, arg.(string))
		}
		cmd := exec.Command(_args[0], _args[1:]...)
		out, err := cmd.Output()
		if err != nil {
			// FIXME: Panic?
			panic(err)
		}
		return string(out)
	}
	panic("Unknown function: " + caller)
}
