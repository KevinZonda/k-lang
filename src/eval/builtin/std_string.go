package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"strconv"
	"strings"
)

type StdStringLib struct {
	b *BuiltIn
}

func NewStdStringLib(b *BuiltIn) *StdStringLib {
	return &StdStringLib{
		b: b,
	}
}

func (s *StdStringLib) FuncCall(caller string, args []any) any {
	switch caller {
	case "len":
		ensureArgsLen(args, 1)
		return len([]rune(args[0].(string)))
	case "fromAsci":
		ensureArgsLen(args, 1)
		return string(args[0].(int))
	case "trim":
		ensureArgsLen(args, 1)
		return strings.TrimSpace(args[0].(string))
	case "trimLeft":
		ensureArgsLen(args, 1)
		return strings.TrimLeft(args[0].(string), " \t\n\r")
	case "trimRight":
		ensureArgsLen(args, 1)
		return strings.TrimRight(args[0].(string), " \t\n\r")
	case "split":
		ensureArgsLen(args, 2)
		sps := strings.Split(args[0].(string), args[1].(string))
		var res []any
		for _, sp := range sps {
			res = append(res, sp)
		}
		return res
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
