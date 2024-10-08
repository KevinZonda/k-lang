package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestLambdaCall(t *testing.T) {
	code := `
foo := fn (x) {
	println(x+11)
}
foo(14)
`
	expected := "25\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestReturnLambda(t *testing.T) {
	code := `
foo := fn () {
	return fn (x) {
        println(x+11)
        return x + 11
    }
}
m := foo()
k := m(17)
`
	expected := "28\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestReturnClosure(t *testing.T) {
	code := `
foo := () => {
    count = 0
    return () => {
        println(count)
        count = count + 1
    }
}
m := foo()
m()
m()
`
	expected := "0\n1\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestReturnClosureLayer2(t *testing.T) {
	code := `
foo := () => {
    count = 0
    {
    	return () => {
    	    println(count)
        	count = count + 1
    	}
    }
}
m := foo()
m()
m()
`
	expected := "0\n1\n"
	tester.GeneralTest(false, t, code, expected)
}

func TestReturnLambdaInLambda(t *testing.T) {
	code := `
x = () => { return () => { println('LAMBDA') } }
x()()
`
	tester.NoPanic(t, code)
}

func TestReturnLambdaInLambda2(t *testing.T) {
	code := `
x = nil
{
    a = 1
    b = 2
    c = 3
    {
        a := 4
        b := 9
        {
            x = () => {
                MEM()
            }
        }
    }
}
x()
`
	tester.NoPanic(t, code)
}
