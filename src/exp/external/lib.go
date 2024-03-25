package external

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/builtin"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/common"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type ExternelLibrary struct {
	l    common.Server
	pack string
}

func NewExternelLibrary(l common.Server, pack string) *ExternelLibrary {
	return &ExternelLibrary{l: l, pack: pack}
}

func (e ExternelLibrary) FuncCall(name string, args []any) obj.ILibraryCall {
	i := e.l.InvokeFunc(name, args...)
	if !i.Success {
		panic("Failed to call function")
	}
	return i
}

func (e ExternelLibrary) GetFunc(name string) *node.FuncBlock {
	fx := builtin.FxToFuncBlock(func(args []any) any {
		i := e.l.InvokeFunc(name, args...)
		if !i.Success {
			panic("Failed to call function")
		}
		return i.Result
	})
	fx.Args[0].Param = true
	fx.Args[0].ParamEmpty = true
	return fx
}

func (e ExternelLibrary) GetObjList() map[string]*obj.Object {
	return nil
}

var _ obj.ILibrary = &ExternelLibrary{}
