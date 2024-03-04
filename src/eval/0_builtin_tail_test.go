package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestBuildInTail(t *testing.T) {
	code := `
x = {"Hi": 18}
println(x.keys())
println(x.values())
x.remove("Hi")
println(x)
println(x.len())
println([1].len())
nil
return
`
	tester.NoPanic(t, code)
}

func TestStrBuildInTail(t *testing.T) {
	code := `
x = ""
x.upper()
x.lower()
x.split("")
x.contains("18")
x.starts_with("22")
x.ends_with("22")
x.len()
`
	tester.NoPanic(t, code)
	code = `
x = 'true'
x.to_bool()`
	tester.NoPanic(t, code)

	code = `
x = '11'
x.to_int()`
	tester.NoPanic(t, code)

	code = `
x = '11'
x.to_num()`
	tester.NoPanic(t, code)
}
