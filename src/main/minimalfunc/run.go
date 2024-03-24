package minimalfunc

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/compressor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/module"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"strings"
)

func Run(input string) {
	if input == "" || input == "." || strings.HasSuffix(input, ".mod") {
		bs, e := iox.ReadAllText(module.DEFAULT_K_MOD)
		panicx.PanicIfNotNil(e, e)
		mod := module.LoadFromText(string(bs))
		input = mod.Entry
		if input == "" {
			input = module.DEFAULT_ENTRY
		}
	}
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
		var errs parserHelper.CodeErrors
		ast, errs = parserHelper.Ast(string(bs))
		errs.PanicIfError()
	}
	ev := eval.New().WithBasePath(input)
	ev.DoSafely(ast).PrintPanic()
}
