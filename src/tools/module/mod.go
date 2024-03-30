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
	opt := hjson.DefaultOptions()
	opt.EmitRootBraces = false
	bs, _ := hjson.MarshalWithOptions(m, opt)
	return string(bs)
}

const DEFAULT_PKG_PLACE = "vendor"
const DEFAULT_ENTRY = "main.k"
const DEFAULT_K_MOD = "k.mod"

func (m Mod) Restore() {
	os.Mkdir(DEFAULT_PKG_PLACE, 0755)
	for _, dep := range m.Dependency {
		d := ParseDep(dep)
		restore(d)
	}

}

type Dependency struct {
	Url        string `json:"url"`
	Branch     string `json:"branch"`
	TargetName string `json:"targetName"`
}

func (d Dependency) String() string {
	u := d.Url
	if d.Branch != "" {
		u = u + "@" + d.Branch
	}
	if d.TargetName != "" {
		u = u + " " + d.TargetName
	}
	return u
}

func ParseDep(s string) Dependency {
	s = strings.TrimSpace(s)
	if s == "" {
		return Dependency{}
	}
	d := Dependency{}
	ds := strings.SplitN(s, " ", 2)
	d.Url = ds[0]
	if len(ds) > 1 {
		d.TargetName = ds[1]
	}
	if strings.Contains(d.Url, "@") {
		dsBranch := strings.SplitN(d.Url, "@", 2)
		d.Url = dsBranch[0]
		if len(dsBranch) > 1 {
			d.Branch = dsBranch[1]
		}
	}
	return d

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

func restore(d Dependency) {
	// run git clone
	bin := "git"
	args := []string{"clone", "--depth", "1"}
	// run git clone
	// run go mod init
	// run go mod tidy
	if d.Branch != "" {
		args = append(args, "--branch", d.Branch)
	}
	args = append(args, normUrl(d.Url))
	if d.TargetName != "" {
		args = append(args, d.TargetName)
	}
	cmd := exec.Command(bin, args...)
	fmt.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Dir = DEFAULT_PKG_PLACE
	err := cmd.Run()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			panic("git not found, please install git first")
		}
		panic(err)
	}
}
