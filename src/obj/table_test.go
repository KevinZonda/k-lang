package obj_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/tx"
	"testing"
)

func TestObjTable(t *testing.T) {
	code := `
fn foo() {
  x := 14
  z()
  println(x)
}

fn z() {
  x := 17
}

foo()
`
	expected := `14
`
	if err := tx.IsStdoutAsExpected(true, func() {
		ast, _ := parserHelper.Ast(code)
		e := eval.New(ast, "")
		e.Do()
	}, expected); err != nil {
		t.Fatal(err)
	}
}
