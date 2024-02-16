package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"strings"
)

var libMap = map[string]obj.ILibrary{}

func GetLibrary(name string) obj.ILibrary {
	if lib, ok := libMap[name]; ok {
		return lib
	}
	if strings.HasPrefix(name, "std/") {
		name = name[4:]
	}
	var lib obj.ILibrary
	switch name {
	case "string":
		lib = NewStdStringLib()
	case "console":
		lib = NewStdConsoleLib()
	case "exec":
		lib = NewStdExecLib()
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
