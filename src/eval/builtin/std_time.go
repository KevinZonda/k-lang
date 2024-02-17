package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"time"
)

type StdTimeLib struct{}

func NewStdTimeLib() *StdTimeLib {
	return &StdTimeLib{}
}

func (s *StdTimeLib) FuncCall(b obj.BuiltInInterface, caller string, args []any) obj.ILibraryCall {
	switch caller {
	case "now":
		ensureArgsLen(args, 0)
		return resultVal(time.Now().String())
	case "unix":
		ensureArgsLen(args, 0)
		return resultVal(time.Now().Unix())
	case "unixNano":
		ensureArgsLen(args, 0)
		return resultVal(time.Now().UnixNano())
	case "sleep":
		ensureArgsLen(args, 1)
		time.Sleep(time.Duration(args[0].(int)) * time.Millisecond)
	}
	panic("Unknown function: " + caller)
}
