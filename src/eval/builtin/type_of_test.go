package builtin_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestTypeOf(t *testing.T) {
	code := `
println(typeOf(1))
`
	p, _ := parserHelper.Ast(code)
	e := eval.New()
	e.Do(p)

}
