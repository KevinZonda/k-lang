package builtin_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestTypeOf(t *testing.T) {
	code := `
open "console"
println(typeOf(1))
println(typeOf([1, 2, 3]))
println(typeOf(true))
console.writeln(typeOf(1.2))
console.write(typeOf(""))
println(len("18"))
println(len([]))
println(len({}))
# FIXME: println(len(map{}))
println(typeOf(map{"11": 11}))
`
	p, _ := parserHelper.Ast(code)
	e := eval.New()
	e.Do(p)

}
