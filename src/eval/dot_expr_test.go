package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestDotExpr(t *testing.T) {
	code := `
a.z.y()
`
	a, _ := parserHelper.Ast(code)
	eval.New(a, "").Do()
}
