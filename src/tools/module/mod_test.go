package module_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/module"
	"os"
	"testing"
)

func TestMod(t *testing.T) {
	initEnvironment()
	defer initEnvironment()
	// t.Skip("skipping test")
	m := module.LoadFromText(`
k: 0.2.26
deps: [
  github.com/hjson/hjson-go v4.4.0
]
`)
	fmt.Println(m)
	m.Restore()
}

func initEnvironment() {
	os.RemoveAll("vendor")
}
