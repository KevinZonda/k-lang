package funcs

import (
	"encoding/json"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
)

func Ast(input string, output string) {
	txt, err := iox.ReadAllText(input)
	panicx.PanicIfNotNil(err, err)

	ast := parserHelper.Ast(txt)
	bs, err := json.MarshalIndent(ast, "", "    ")
	panicx.PanicIfNotNil(err, err)

	err = iox.WriteAllText(output, string(bs))
	panicx.PanicIfNotNil(err, err)
}
