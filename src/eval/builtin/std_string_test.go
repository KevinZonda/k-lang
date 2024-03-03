package builtin_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestStdStringLib(t *testing.T) {

	code := `
open "std/string" as s
x := s.int("128")
println(typeOf(x))
println(x)
z := s.float("128.0")
println(typeOf(z))
println(z)
z := s.trimLeft("  128.0    ")
println(z)
z := s.trimRight("  128.0    ")
println(z)
z := s.trim("  128.0    ")
println(z)
println(s.split("1 2",' '))
println(s.fromAscii(67))
println(s.len("Hi"))
`
	expected := `int
128
num
128
128.0    
  128.0
128.0
[1 2]
C
2
`

	tester.GeneralTest(false, t, code, expected)
}
