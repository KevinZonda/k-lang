package parserHelper_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestErrOutput(t *testing.T) {
	code := `
if x == 17 {
}`
	p := parserHelper.FromString(code)
	_ = p.Ast()
	for _, err := range p.Errors() {
		fmt.Println(err)
	}
}
