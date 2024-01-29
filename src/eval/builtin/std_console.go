package builtin

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type StdConsoleLib struct{}

func NewStdConsoleLib() *StdConsoleLib {
	return &StdConsoleLib{}
}

func (c *StdConsoleLib) FuncCall(b obj.BuiltInInterface, caller string, args []any) any {
	switch caller {
	case "readln":
		ensureArgsLen(args, 0)
		var s string
		_, err := fmt.Fscanln(b.GetStdin(), &s)
		if err != nil {
			panic(err)
		}
		return s
	case "write":
		_b := b.(BuiltIn)
		_b.Print(args...)
		return nil
	case "writeln":
		_b := b.(BuiltIn)
		_b.Println(args...)
		return nil
	}
	panic("Unknown function: " + caller)
}

var _ obj.ILibrary = (*StdConsoleLib)(nil)
