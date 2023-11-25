package builtin

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/objType"

var libMap = map[string]objType.ILibrary{}

type ITypeOf interface {
	TypeOf() string
}

func GetLibrary(name string) objType.ILibrary {
	if lib, ok := libMap[name]; ok {
		return lib
	}
	var lib objType.ILibrary
	switch name {
	case "std/string", "string":
		lib = NewStdStringLib()
	case "std/console", "console":
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
