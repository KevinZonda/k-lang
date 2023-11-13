package pruner_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/pruner"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestBinaryOper(t *testing.T) {
	v, _ := parserHelper.Ast("7+1+2+3+a")
	t.Log(v)
	fmt.Println("------------------")
	v = pruner.Prune(v)
	t.Log(v)
}
