package funcs

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/compressor"
	"github.com/KevinZonda/GoX/pkg/iox"
	"path/filepath"
	"strings"
)

func Compile(input string, output string) {
	if output == "" {
		output = strings.TrimSuffix(input, filepath.Ext(input)) + ".ast"
	}

	str, e := iox.ReadAllText(input)
	if e != nil {
		panic("Error: cannot read file:" + input)
	}

	ast, errs := parserHelper.Ast(str)
	errs.PanicIfError()

	c, ce := compressor.Compress(ast)
	if ce != nil {
		panic(ce)
	}

	e = iox.WriteAllBytes(output, c)
	if e != nil {
		panic("Error: cannot write to file:" + output)
	}
}
