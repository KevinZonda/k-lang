package builtin

type ILibrary interface {
	FuncCall(name string, args []any) any
}

var libMap = map[string]ILibrary{}

func GetLibrary(name string) ILibrary {
	if lib, ok := libMap[name]; ok {
		return lib
	}
	var lib ILibrary
	switch name {
	case "string":
		lib = NewStdStringLib()
		libMap[name] = lib
		return lib
	}
	return nil
}
