package builtin

import (
	"bufio"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type StdConsoleLib struct {
	FBLibrary
}

func NewStdConsoleLib(b obj.StdIO) *StdConsoleLib {
	l := FBLibrary{
		IODep: true,
		IO:    b,
	}
	printF := FxToFuncBlock(func(a []any) {
		Print(l.IO, a)
	})
	printF.Args[0].Param = true

	printLnF := FxToFuncBlock(func(a []any) {
		Println(l.IO, a)
	})
	printLnF.Args[0].Param = true

	l.V = map[string]*node.FuncBlock{
		"readln": FxToFuncBlock(func() string {
			reader := bufio.NewReader(l.IO.GetStdin())
			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			return text
		}),
		"write":   printF,
		"writeln": printLnF,
	}
	return &StdConsoleLib{
		FBLibrary: l,
	}
}

func (c *StdConsoleLib) GetObjList() map[string]*obj.Object {
	return nil
}

var _ obj.ILibrary = (*StdConsoleLib)(nil)
