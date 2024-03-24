package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestRefInt(t *testing.T) {
	code := `
fn add1(&x) {
  	x = x + 1
}
x := 12
println(x)
add1(x)
println(x)
`
	expected := `12
13
`
	tester.GeneralTest(false, t, code, expected)
}

func TestRefArr(t *testing.T) {
	code := `
fn add1(&x) {
  	x = x + 1
}
x := []
println(x)
add1(x)
println(x)
`
	expected := `[]
[1]
`
	tester.GeneralTest(false, t, code, expected)
}

func TestRefStruct(t *testing.T) {
	code := `
struct color {
  int r, g, b
  int a = 20
}
fn add1(&x) {
  	x.g = 256
}
x := color{}
println(x)
add1(x)
println(x)
`
	expected := `color {r: 0, g: 0, b: 0, a: 20}
color {r: 0, g: 256, b: 0, a: 20}
`
	tester.GeneralTest(false, t, code, expected)
}

func TestRefBinOper(t *testing.T) {
	code := `
x = 1
y = &x
z = y + 1
println(z)
`
	expected := `2
`
	tester.GeneralTest(false, t, code, expected)
}

func TestDoubleRefUOper(t *testing.T) {
	code := `
x = 1
y =& &x
println(-y)
`
	expected := `-1
`
	tester.GeneralTest(false, t, code, expected)
}

func TestRefDel(t *testing.T) {
	code := `
x = 1
y = &x
y = 19
println(y)
`
	expected := `19
`
	tester.GeneralTest(false, t, code, expected)
}

func TestRefLambdaCall(t *testing.T) {
	code := `
x = (y) => {
    println(y + 1)
}
z = &x
z(18)
`
	expected := `19
`
	tester.GeneralTest(false, t, code, expected)
}

func TestRefLambdaCall2(t *testing.T) {
	code := `
x = (y) => {
    println(y + 1)
}
a = 18
x(&a)
`
	expected := `19
`
	tester.GeneralTest(false, t, code, expected)
}

//func TestRefIgnore(t *testing.T) {
//	code := `
//x = 12
//y := &x
//y = 13
//println(x)
//println(y)
//`
//	expected := `12
//13
//`
//	tester.GeneralTest(false, t, code, expected)
//}
