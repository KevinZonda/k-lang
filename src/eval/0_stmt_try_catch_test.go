package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestTryCatchStmt(t *testing.T) {
	code := `
try {
  panic("18")
} catch {
  println("Caught")
}
`
	expected := "Caught\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestTryCatch2Stmt(t *testing.T) {
	code := `
try {
  panic("18")
} catch(v){
  println(v)
}
`
	expected := "18\n"
	tester.GeneralTest(false, t, code, expected)
}
