package builtin

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/objType"
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
		Print(args...)
		return nil
	case "writeln":
		Println(args...)
		return nil
	}
	panic("Unknown function: " + caller)
}

var _ objType.ILibrary = (*StdConsoleLib)(nil)
