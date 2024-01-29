package builtin

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type StdConsoleLib struct {
	b *BuiltIn
}

func NewStdConsoleLib(b *BuiltIn) *StdConsoleLib {
	return &StdConsoleLib{b: b}
}

func (c *StdConsoleLib) FuncCall(caller string, args []any) any {
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
		c.b.Print(args...)
		return nil
	case "writeln":
		c.b.Println(args...)
		return nil
	}
	panic("Unknown function: " + caller)
}

var _ obj.ILibrary = (*StdConsoleLib)(nil)
