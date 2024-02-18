package buildconst

import (
	"fmt"
	"runtime"
	"strings"
)

const LanguageVersion = "6.6.6"

var (
	BuildTime   string = "2000/01/01"
	CommitHash  string = "000000"
	BuildBranch string = "main"
)

func init() {
	if BuildBranch == "" {
		BuildBranch = "main"
	}
	if CommitHash == "" {
		CommitHash = "000000"
	}
}

func Msg() string {
	return fmt.Sprintf("K Language %s (%s|%s, %s) [%s %s (%s)] on %s",
		LanguageVersion, CommitHash, BuildBranch, BuildTime,
		runtime.Version(), runtime.Compiler, strings.ToUpper(runtime.GOARCH), runtime.GOOS)
}
