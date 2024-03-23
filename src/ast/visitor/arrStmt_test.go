package visitor_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/jout"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestArr(t *testing.T) {
	ast, _ := parserHelper.Ast("[1, 2, 3, 4]")
	jout.Println(ast)
}
