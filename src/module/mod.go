package module

import (
	"errors"
	"fmt"
	"github.com/hjson/hjson-go/v4"
	"os"
	"os/exec"
	"strings"
)

type Mod struct {
	KLangVersion string   `json:"k"`
	Entry        string   `json:"entry,omitempty"`
	Dependency   []string `json:"deps"`
}

func (m Mod) String() string {
	bs, _ := hjson.MarshalWithOptions(m, hjson.DefaultOptions())
	return string(bs)
}

const DEFAULT_PKG_PLACE = "vendor"
const DEFAULT_ENTRY = "main.k"
const DEFAULT_K_MOD = "k.mod"

func (m Mod) Restore() {
	os.Mkdir(DEFAULT_PKG_PLACE, 0755)
	for _, dep := range m.Dependency {
		d := strings.TrimSpace(dep)
		if d == "" {
			continue
		}
		ds := strings.SplitN(d, " ", 2)
		if len(ds) == 1 || len(ds) > 2 {
			panic("invalid dependency")
		}
		restore(ds[0], ds[1])
	}

}

func normUrl(url string) string {
	if strings.HasPrefix(url, "https://") {
		return url
	}
	if strings.HasPrefix(url, "http://") {
		return url
	}
	return "https://" + url
}

func restore(url string, tag string) {
	// run git clone
	bin := "git"
	args := []string{"clone", "--depth", "1", "--branch", tag, normUrl(url)}
	// run git clone
	// run go mod init
	// run go mod tidy
	cmd := exec.Command(bin, args...)
	fmt.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Dir = DEFAULT_PKG_PLACE
	err := cmd.Run()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			panic(err)
		}
		panic(err)
	}
}
