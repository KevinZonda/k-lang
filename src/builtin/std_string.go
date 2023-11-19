package builtin

import (
	"strings"
)

type StdStringLib struct {
}

func NewStdStringLib() *StdStringLib {
	return &StdStringLib{}
}

func (s *StdStringLib) FuncCall(caller string, args []any) any {
	switch caller {
	case "len":
		return len([]rune(args[0].(string)))
	case "trim":
		return strings.TrimSpace(args[0].(string))
	}
	panic("Unknown function: " + caller)
}

var _ ILibrary = (*StdStringLib)(nil)
