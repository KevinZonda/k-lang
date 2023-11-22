package builtin

import (
	"fmt"
)

type StdConsoleLib struct{}

func NewStdConsoleLib() *StdConsoleLib {
	return &StdConsoleLib{}
}

func (s *StdConsoleLib) FuncCall(caller string, args []any) any {
	switch caller {
	case "readln":
		ensureArgsLen(args, 0)
		var s string
		_, err := fmt.Scanln(&s)
		if err != nil {
			panic(err)
		}
		return s
	case "write":
		for _, arg := range args {
			fmt.Print(arg)
		}
		return nil
	case "writeln":
		for _, arg := range args {
			fmt.Print(arg)
		}
		fmt.Println()
		return nil
	}
	panic("Unknown function: " + caller)
}

var _ ILibrary = (*StdConsoleLib)(nil)