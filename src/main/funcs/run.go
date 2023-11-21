package funcs

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/compressor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
)

func Run(input string) {
	bs, e := iox.ReadAllByte(input)

	panicx.PanicIfNotNil(e, e)
	var ast tree.Ast

	if compressor.IsCompressed(bs) {
		var ce compressor.CompressorError
		ast, ce = compressor.Decompress(bs)
		if ce != nil {
			panic(ce)
		}
	} else {
		var errs []parserHelper.CodeError
		ast, errs = parserHelper.Ast(string(bs))
		if len(errs) > 0 {
			parserHelper.PrintAllCodeErrors(errs)
			panic("Parse failed.")
		}
	}
	ev := eval.New(ast, input)
	ev.Do()
}
