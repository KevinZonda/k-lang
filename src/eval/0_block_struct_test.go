package eval_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/tester"
	"testing"
)

func TestEvalStructLiteral(t *testing.T) {
	src := `
x := struct {
   red: 10,
   green: 20,
   blue: 30,
}
println(x)
	`
	tester.GeneralTest(true, t, src, "struct {red: 10, green: 20, blue: 30}\n")
}

func TestEvalStructLiteral2(t *testing.T) {
	src := `
struct foo {
  int x = 12
  bool y = true
}
x := foo{}
println(x)
	`
	tester.GeneralTest(true, t, src, "foo {x: 12, y: true}\n")
}

func TestEvalStructLiteral3(t *testing.T) {
	src := `
struct foo {
  int x = 12
  bool y = true
  help = fn () {
    println("hello")
    println(self.x)
  }
}
x := foo{}
x.help()
	`
	tester.GeneralTest(true, t, src, "hello\n12\n")
}
