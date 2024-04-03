package visualizer_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/visualizer"
	"testing"
)

func TestV(t *testing.T) {
	code := `
struct color {
  fn foo() {}
}
struct ll {
	next: ll?
    val: int
}
x := ll {val: 1}
x.next = ll {val: 2}
y = ll{val: 3}
x.next.next = &y
y.val = struct {
   a: 18
   b : "Hi"
   c : color{}
   d : [1, 2, 3]
}
`
	e := eval.New()

	ast, _ := parserHelper.Ast(code)
	e.DoSafely(ast)
	m := e.GetMemory()
	obj2, _ := m.Bottom().Get("x")
	j := obj.TreeAnyW("x", obj2, false)
	// jout(j)

	fmt.Println(j.String())

	k := visualizer.Visualize("x", obj2)
	fmt.Println(k.StringIdent(0))

}
