package eval_test

import "testing"

func TestStdString(t *testing.T) {
	code := `
open "string" as str
println(str.len("  hello  "))
println(str.trim("  hello  "))
`
	GeneralTest(true, t, code, "9\nhello\n")
}
