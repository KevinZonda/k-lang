package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"strings"
)

var cache = map[string]library{}

func GetLibrary(name string, io obj.StdIO) obj.ILibrary {
	if lib, ok := cache[name]; ok {
		return lib
	}
	if strings.HasPrefix(name, "std/") {
		name = name[4:]
	}
	var lib library
	switch name {
	case "string":
		lib = NewStdStringLib()
	case "console":
		lib = NewStdConsoleLib(io)
	case "exec":
		lib = NewStdExecLib()
	case "time":
		lib = NewStdTimeLib()
	case "math":
		lib = NewStdMathLib()
	}
	if lib != nil && !lib.IsIODep() {
		cache[name] = lib
	}
	return lib
}

type library interface {
	obj.ILibrary
	IsIODep() bool
}
