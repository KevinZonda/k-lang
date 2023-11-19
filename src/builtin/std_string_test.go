package builtin_test

import (
	"testing"
)

func TestStdStringLib(t *testing.T) {

	code := `
open "string" as s
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
`
	expected := `int
128
float64
128
128.0    
  128.0
`

	GeneralTest(false, t, code, expected)
}
