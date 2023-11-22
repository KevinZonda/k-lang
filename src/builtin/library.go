package builtin

type ILibrary interface {
	FuncCall(name string, args []any) any
}

var libMap = map[string]ILibrary{}

type ITypeOf interface {
	TypeOf() string
}

func GetLibrary(name string) ILibrary {
	if lib, ok := libMap[name]; ok {
		return lib
	}
	var lib ILibrary
	switch name {
	case "string":
		lib = NewStdStringLib()
	case "console":
		lib = NewStdConsoleLib()
	}
	if lib != nil {
		libMap[name] = lib
	}
	return lib
}

func ensureArgsLen(args []any, n int) {
	if len(args) != n {
		panic("Wrong number of arguments")
	}
}
