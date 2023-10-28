package main

import (
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/parserHelper"
)

func main() {
	p := parserHelper.FromString("1 + 2 + 3")
	p.Ast()
}
