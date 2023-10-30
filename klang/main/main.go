package main

import (
	"encoding/json"
	"fmt"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/eval"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/parserHelper"
)

func main() {
	p := parserHelper.FromString("(1 + 2.1) * (3+4)+\"114514\"\n struct x { int a }")
	ast := p.Ast()
	for idx, node := range ast {
		bs, _ := json.MarshalIndent(node, "", "    ")
		fmt.Println(idx, "->", string(bs))
	}
	e := eval.New(ast)
	e.Do()
	fmt.Println("Evaluated as", e.It())
}
