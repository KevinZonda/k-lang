package builtin_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestStdMathLib(t *testing.T) {

	code := `
open "std/math" as m
assert = (b) => {
	if !b {
		panic("assertion failed")
	}
}

assert(m.pi == 3.141592653589793)
assert(m.e == 2.718281828459045)
assert(m.abs(-1) == 1)
assert(m.pow(2, 3) == 8)
assert(m.sqrt(9) == 3)
assert(m.cbrt(27) == 3)
assert(m.sin(0) == 0)
assert(m.cos(0) == 1)
assert(m.tan(0) == 0)
assert(m.asin(0) == 0)
assert(m.acos(1) == 0)
assert(m.atan(0) == 0)
assert(m.exp(0) == 1)
assert(m.ln(1) == 0)
assert(m.log(1) == 0)
assert(m.log10(1) == 0)
assert(m.log2(1) == 0)
assert(m.randInt() != m.randInt())
assert(m.randAlphabet() != m.randAlphabet())
x = m.randIntN(10)
assert(x < 10)
x = m.randNum()
assert(x >= 0 && x <= 1)
assert(m.getRandSeed() == 0)
m.setRandSeed(11)
assert(m.getRandSeed() == 11)
m.randSeedTime()
assert(m.getRandSeed() != 0)
m.resetRandSeed()
assert(m.getRandSeed() == 0)
x = [1,2,3,4]
m.shuffle(x)
assert(x != [1,2,3,4])
m.randSeed(19)
assert(m.randInt() == m.randInt())
assert(m.randIntN(10) == m.randIntN(10))
assert(m.randNum() == m.randNum())
assert(m.randAlphabet() == m.randAlphabet())
x = [1,2,3,4]
m.shuffle(x)
assert(x != [1,2,3,4])
`

	tester.NoPanic(t, code)
}
