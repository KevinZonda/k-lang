package buildconst

import (
	"fmt"
	"os"
)

func Exit(code int) {
	if TestMod {
		panic(fmt.Sprint(code))
	}
	os.Exit(code)
}

var TestMod = false
