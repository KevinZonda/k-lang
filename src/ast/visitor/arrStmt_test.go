package visitor_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/jout"
	"testing"
)

func TestArr(t *testing.T) {
	ast, _ := parserHelper.Ast("[1, 2, 3, 4]")
	jout.Println(ast)
}
