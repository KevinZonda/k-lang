package funcs

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/jout"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	compressor2 "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/compressor"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
)

func ShowAst(input string) {
	bs, e := iox.ReadAllByte(input)
	panicx.PanicIfNotNil(e, e)
	var ast tree.Ast
	if compressor2.IsCompressed(bs) {
		var ce compressor2.CompressorError
		ast, ce = compressor2.Decompress(bs)
		if ce != nil {
			panic(ce)
		}
	} else {
		var errs parserHelper.CodeErrors
		ast, errs = parserHelper.Ast(string(bs))
		errs.PanicIfError()
	}
	jout.Println(ast)
}
