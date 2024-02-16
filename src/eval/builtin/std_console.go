package builtin

import (
	"bufio"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type StdConsoleLib struct{}

func NewStdConsoleLib() *StdConsoleLib {
	return &StdConsoleLib{}
}

func (c *StdConsoleLib) FuncCall(b obj.BuiltInInterface, caller string, args []any) obj.ILibraryCall {
	switch caller {
	case "readln":
		ensureArgsLen(args, 0)
		reader := bufio.NewReader(b.GetStdin())
		text, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}
		return resultVal(text)
	case "write":
		_b := b.(BuiltIn)
		_b.Print(args...)
		return resultNoVal()
	case "writeln":
		_b := b.(BuiltIn)
		_b.Println(args...)
		return resultNoVal()
	}
	panic("Unknown function: " + caller)
}

var _ obj.ILibrary = (*StdConsoleLib)(nil)
