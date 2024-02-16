package buildconst

import (
	"fmt"
	"runtime"
	"strings"
)

const LanguageVersion = "1.0.0"

var (
	BuildTime  string = "2000/01/01"
	CommitHash string = "000000"
)

func Msg() string {
	return fmt.Sprintf("K Language %s (%s, %s) [%s %s (%s)] on %s",
		LanguageVersion, CommitHash, BuildTime,
		runtime.Version(), runtime.Compiler, strings.ToUpper(runtime.GOARCH), runtime.GOOS)
}
