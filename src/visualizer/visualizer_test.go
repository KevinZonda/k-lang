package visualizer_test

import (
	"encoding/json"
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/visualizer"
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
   v : color{}
}
`
	e := eval.New()

	ast, _ := parserHelper.Ast(code)
	e.DoSafely(ast)
	m := e.GetMemory()
	obj, _ := m.Bottom().Get("x")
	j := visualizer.TreeAny(nil, "x", obj)
	// jout(j)

	fmt.Println(j.String())

}

func jout(v any) {
	bs, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bs))

}
