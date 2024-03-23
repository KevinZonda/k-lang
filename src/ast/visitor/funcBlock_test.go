package visitor_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/jout"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestFuncBlock(t *testing.T) {
	ast, _ := parserHelper.Ast(`
fn add(int x, int y) float {
    y := x + 1
}
`)
	jout.Println(ast)
	ast, _ = parserHelper.Ast(`
fn add() float {
    y := x + 1
}
`)
	jout.Println(ast)
	ast, _ = parserHelper.Ast(`
fn add() {
    y := x + 1
}
`)
	jout.Println(ast)
}

func TestFuncCall(t *testing.T) {
	ast, _ := parserHelper.Ast(`
add()`)
	jout.Println(ast)
	ast, _ = parserHelper.Ast(`
add(1, 2)`)
	jout.Println(ast)
}
