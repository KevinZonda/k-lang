package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"strconv"
	"strings"
)

type StdStringLib struct{}

func NewStdStringLib() *StdStringLib {
	return &StdStringLib{}
}

func (s *StdStringLib) FuncCall(caller string, args []any) any {
	switch caller {
	case "len":
		ensureArgsLen(args, 1)
		return len([]rune(args[0].(string)))
	case "trim":
		ensureArgsLen(args, 1)
		return strings.TrimSpace(args[0].(string))
	case "trimLeft":
		ensureArgsLen(args, 1)
		return strings.TrimLeft(args[0].(string), " \t\n\r")
	case "trimRight":
		ensureArgsLen(args, 1)
		return strings.TrimRight(args[0].(string), " \t\n\r")
	case "int":
		ensureArgsLen(args, 1)
		i, e := strconv.ParseInt(args[0].(string), 10, 64)
		if e != nil {
			panic(e)
		}
		return int(i)
	case "float":
		ensureArgsLen(args, 1)
		f, e := strconv.ParseFloat(args[0].(string), 64)
		if e != nil {
			panic(e)
		}
		return f
	}
	panic("Unknown function: " + caller)
}

var _ obj.ILibrary = (*StdStringLib)(nil)
