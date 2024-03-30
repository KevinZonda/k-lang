package funcs

import (
	"encoding/json"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
)

func Ast(input string, output string) {
	if output == "" {
		ShowAst(input)
		return
	}
	txt, err := iox.ReadAllText(input)
	if err != nil {
		panic("Error: cannot read file:" + input)
	}

	ast, errs := parserHelper.Ast(txt)
	errs.PanicIfError()

	bs, _ := json.MarshalIndent(ast, "", "    ")

	err = iox.WriteAllText(output, string(bs))
	if err != nil {
		panic("Error: cannot write to file:" + output)
	}
}
