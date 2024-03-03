package builtin_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestStdTimeLib(t *testing.T) {

	code := `
open "std/time" as s
println(s.now())
println(s.unix())
println(s.unixNano())
s.sleep(10)


`
	tester.NoPanic(t, code)
}
