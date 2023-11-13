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
		ast, _ = compressor.Decompress(bs)
	} else {
		ast, _ = parserHelper.Ast(string(bs))
	}
	ev := eval.New(ast)
	ev.Do()
}
