package builtin

import (
	"errors"
	"os/exec"

	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type StdExecLib struct{}

func NewStdExecLib() *StdExecLib {
	return &StdExecLib{}
}
func (s *StdExecLib) FuncCall(b obj.StdIO, caller string, args []any) obj.ILibraryCall {
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
			// Q
			var exitError *exec.ExitError
			if errors.As(err, &exitError) {
				// log.Println(exitError.ExitCode())
			}
			panic(err)
		}
		return resultVal(string(out))
	}
	panic("Unknown function: " + caller)
}
