package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestStdString(t *testing.T) {
	code := `
open "string" as str
println(str.len("  hello  "))
println(str.trim("  hello  "))
`
	tester.GeneralTest(true, t, code, "9\nhello\n")
}
