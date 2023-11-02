package visitor_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/jout"
	"testing"
)

func TestFuncBlock(t *testing.T) {
	ast := parserHelper.Ast(`
fn add(int x, int y) float {
    y := x + 1
}
`)
	jout.Println(ast)
	ast = parserHelper.Ast(`
fn add() float {
    y := x + 1
}
`)
	jout.Println(ast)
	ast = parserHelper.Ast(`
fn add() {
    y := x + 1
}
`)
	jout.Println(ast)
}
