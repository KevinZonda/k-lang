package main

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/symbol"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/jout"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
)

func main() {
	ast, err := parserHelper.Ast(`
func main() {
	println("Hello, World!")
}

func add(x : int, y : int) : int {
	return x + y
}

func min(x, y : int) : int {
	return x + y
}

x = 18

struct Point {
	x : int
	y : int
	z = 18
}

_y = 19
`)
	err.PanicIfError()
	f := symbol.AnalyseAst(ast)
	jout.Println(f)
}
