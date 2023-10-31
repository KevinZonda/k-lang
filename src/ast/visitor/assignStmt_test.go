package visitor_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"testing"
)

func TestAssignStmt(t *testing.T) {
	parserHelper.Ast("int a[11][12][17].x.y := 1")
}
