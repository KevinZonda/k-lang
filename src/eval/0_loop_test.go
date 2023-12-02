package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tester"
	"testing"
)

func TestNestedLoop(ts *testing.T) {
	code := `
for (i := 0; i <= 5; i := i + 1) {
    if (i == 4) {
        continue
    }
    for (j := i; j <= 5; j := j + 1) {
        print(i + ", " + (j + 1) + " | ")
    }
    println()
}`
	expected := `0, 1 | 0, 2 | 0, 3 | 0, 4 | 0, 5 | 0, 6 | 
1, 2 | 1, 3 | 1, 4 | 1, 5 | 1, 6 | 
2, 3 | 2, 4 | 2, 5 | 2, 6 | 
3, 4 | 3, 5 | 3, 6 | 
5, 6 | 
`
	tester.GeneralTest(true, ts, code, expected)
}

func TestNestedLoop2(ts *testing.T) {
	code := `
for (i := 0; i <= 5; i := i + 1) {
    if (i == 4) {
        break
    }
    for (j := i; j <= 5; j := j + 1) {
        print(i + ", " + (j + 1) + " | ")
    }
    println()
}`
	expected := `0, 1 | 0, 2 | 0, 3 | 0, 4 | 0, 5 | 0, 6 | 
1, 2 | 1, 3 | 1, 4 | 1, 5 | 1, 6 | 
2, 3 | 2, 4 | 2, 5 | 2, 6 | 
3, 4 | 3, 5 | 3, 6 | 
`
	tester.GeneralTest(true, ts, code, expected)
}

func TestWhileLoop(ts *testing.T) {
	code := `
i := 0
for (i <= 5) {
    i := i - (-1)
    if (i == 3) {
        break
    }
    println(i)
}`
	expected := `1
2
`
	tester.GeneralTest(true, ts, code, expected)
}

func TestWhileLoop2(ts *testing.T) {
	code := `
i := 0
for (i <= 5) {
    i := i - (-1)
    if (i == 3) {
        continue
    }
    println(i)
}`
	expected := `1
2
4
5
6
`
	tester.GeneralTest(true, ts, code, expected)
}

func TestIterLoop(ts *testing.T) {
	code := `
for (x : [1, 7, 8]) {
    println(x)
}`
	expected := `1
7
8
`
	tester.GeneralTest(true, ts, code, expected)
}

func TestIterLoop2(ts *testing.T) {
	code := `
for (x : map{"x": 1, "y": 7, "z": 8}) {
    println(x)
}`
	expected := `struct {key: x, val: 1}
struct {key: y, val: 7}
struct {key: z, val: 8}
`
	tester.GeneralTest(true, ts, code, expected)
}
