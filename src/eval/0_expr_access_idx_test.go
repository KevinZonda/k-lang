package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestAccessByIndex(t *testing.T) {
	code := `
x := [12, 18, 19]
println(x[0])
`
	expected := `12
`
	tester.GeneralTest(false, t, code, expected)
}

func TestMapAccessByIndex(t *testing.T) {
	code := `
x := {"a": 12, "z": 18, "n": 19}
println(x["a"])
`
	expected := `12
`
	tester.GeneralTest(false, t, code, expected)
}

func TestAccessByIndexCJK(t *testing.T) {
	code := `
x := "中文测试"
println(x[0])
`
	expected := `中
`
	tester.GeneralTest(false, t, code, expected)
}

func TestAccessByIndexCJK2(t *testing.T) {
	code := `
x := "中文测试"
println(x[0:2])
`
	expected := `中文
`
	tester.GeneralTest(false, t, code, expected)
}

func TestAccessByIndexCJK3(t *testing.T) {
	code := `
x := "中文测试"
println(x[2:])
`
	expected := `测试
`
	tester.GeneralTest(false, t, code, expected)
}

func TestAccessByIndexCJK4(t *testing.T) {
	code := `
x := "中文测试"
println(x[2:0])
`
	expected := `文中
`
	tester.GeneralTest(false, t, code, expected)
}
