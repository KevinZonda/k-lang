package builtin

type ILibrary interface {
	FuncCall(name string, args []any) any
}
