package buildconst

import (
	"fmt"
	"runtime"
)

const LanguageVersion = "1.0.0"

var (
	BuildTime  string = "2000/01/01"
	CommitHash string = "000000"
)

func Msg() string {
	return fmt.Sprintf("K Language %s (%s, %s) %s on %s %s",
		LanguageVersion, CommitHash, BuildTime,
		runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
