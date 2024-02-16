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

func (s *StdStringLib) FuncCall(b obj.BuiltInInterface, caller string, args []any) obj.ILibraryCall {
	switch caller {
	case "len":
		ensureArgsLen(args, 1)
		return resultVal(len([]rune(args[0].(string))))
	case "fromAscii":
		ensureArgsLen(args, 1)
		return resultVal(string(args[0].(rune)))
	case "trim":
		ensureArgsLen(args, 1)
		return resultVal(strings.TrimSpace(args[0].(string)))
	case "trimLeft":
		ensureArgsLen(args, 1)
		return resultVal(strings.TrimLeft(args[0].(string), " \t\n\r"))
	case "trimRight":
		ensureArgsLen(args, 1)
		return resultVal(strings.TrimRight(args[0].(string), " \t\n\r"))
	case "split":
		ensureArgsLen(args, 2)
		sps := strings.Split(args[0].(string), args[1].(string))
		var res []any
		for _, sp := range sps {
			res = append(res, sp)
		}
		return resultVal(res)
	case "int":
		ensureArgsLen(args, 1)
		i, e := strconv.ParseInt(args[0].(string), 10, 64)
		if e != nil {
			panic(e)
		}
		return resultVal(int(i))
	case "float":
		ensureArgsLen(args, 1)
		f, e := strconv.ParseFloat(args[0].(string), 64)
		if e != nil {
			panic(e)
		}
		return resultVal(f)
	}
	panic("Unknown function: " + caller)
}

var _ obj.ILibrary = (*StdStringLib)(nil)
