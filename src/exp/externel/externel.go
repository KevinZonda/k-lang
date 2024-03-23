package externel

import (
	"encoding/json"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

type Server interface {
	AvailablePackage() []string
	PackageInfo() map[string]PackageInfoElement
	InvokeFunc(method string, args ...interface{}) InvokeResult
}

type PackageInfoElement struct {
	Type ElementType `json:"type"`
	Args []ArgType   `json:"args"`
}

type ElementType string

const (
	Function ElementType = "func"
	Variable ElementType = "var"
)

type InvokeResult struct {
	Success   bool `json:"success"`
	HasResult bool `json:"hasResult"`
	Result    any  `json:"result"`
}

func (i InvokeResult) BS() []byte {
	bs, _ := json.Marshal(i)
	return bs
}

func (i InvokeResult) HasValue() bool {
	return i.HasResult
}

func (i InvokeResult) Value() any {
	return i.Result
}

type ExternelLibrary struct {
	l    Server
	pack string
}

func NewExternelLibrary(l Server, pack string) *ExternelLibrary {
	return &ExternelLibrary{l: l, pack: pack}
}

func (e ExternelLibrary) FuncCall(b obj.StdIO, name string, args []any) obj.ILibraryCall {
	i := e.l.InvokeFunc(name, args...)
	if !i.Success {
		panic("Failed to call function")
	}
	return i
}

func (e ExternelLibrary) GetObjList() map[string]*obj.Object {
	return nil
}

var _ obj.ILibrary = &ExternelLibrary{}
