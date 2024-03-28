package external

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/httpExternal"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"strings"
)

const EXTERNAL = true

func GetExternalLibrary(lib string) obj.ILibrary {
	if !EXTERNAL {
		return nil
	}
	if strings.HasPrefix(lib, "https://") || strings.HasPrefix(lib, "http://") {
		return NewExternalLibrary(httpExternal.NewLibrary(lib), "")
	}
	return nil
}
