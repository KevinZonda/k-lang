package external

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternal"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"strings"
)

func GetExternalLibrary(lib string) obj.ILibrary {
	if strings.HasPrefix(lib, "https://") || strings.HasPrefix(lib, "http://") {
		return NewExternelLibrary(httpExternal.NewLibrary(lib), "")
	}
	return nil
}
