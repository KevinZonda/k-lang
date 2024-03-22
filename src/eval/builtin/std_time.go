package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"time"
)

type StdTimeLib struct{}

func NewStdTimeLib() *StdTimeLib {
	return &StdTimeLib{}
}

func (c *StdTimeLib) GetObjList() map[string]*obj.Object {
	return nil
}

func (s *StdTimeLib) FuncCall(b obj.StdIO, caller string, args []any) obj.ILibraryCall {
	switch caller {
	case "now":
		ensureArgsLen(args, 0)
		return resultVal(time.Now().String())
	case "unix":
		ensureArgsLen(args, 0)
		return resultVal(int(time.Now().Unix()))
	case "unixNano":
		ensureArgsLen(args, 0)
		return resultVal(int(time.Now().UnixNano()))
	case "sleep":
		ensureArgsLen(args, 1)
		time.Sleep(time.Duration(args[0].(int)) * time.Millisecond)
		return resultNoVal()
	}
	panic("Unknown function: " + caller)
}
