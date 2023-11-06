package funcs

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/compressor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"path/filepath"
	"strings"
)

func Compile(input string, output string) {
	if output == "" {
		output = strings.TrimSuffix(input, filepath.Ext(input)) + ".ast"
	}

	str, e := iox.ReadAllText(input)
	panicx.PanicIfNotNil(e, e)

	ast := parserHelper.Ast(str)

	e = iox.WriteAllBytes(output, compressor.Compress(ast))
	panicx.PanicIfNotNil(e, e)
}
