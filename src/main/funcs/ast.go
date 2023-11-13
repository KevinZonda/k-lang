package funcs

import (
	"encoding/json"
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
)

func Ast(input string, output string) {
	txt, err := iox.ReadAllText(input)
	if err != nil {
		panic("Error: cannot read file:" + input)
	}

	ast, errs := parserHelper.Ast(txt)
	if len(errs) >= 0 {
		printAllCodeErros(errs)
		panic("Parse failed.")
	}
	bs, _ := json.MarshalIndent(ast, "", "    ")

	err = iox.WriteAllText(output, string(bs))
	if err != nil {
		panic("Error: cannot write to file:" + output)
	}
}

func printAllCodeErros(errs []parserHelper.CodeError) {
	fmt.Println("Parsing Error(s):")
	for i, e := range errs {
		fmt.Println("[", i, "] :", e)
	}
}
