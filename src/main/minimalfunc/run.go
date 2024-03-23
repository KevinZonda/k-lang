package minimalfunc

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	compressor2 "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/compressor"
	module2 "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/module"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"strings"
)

func Run(input string) {
	if input == "" || input == "." || strings.HasSuffix(input, ".mod") {
		bs, e := iox.ReadAllText(module2.DEFAULT_K_MOD)
		panicx.PanicIfNotNil(e, e)
		mod := module2.LoadFromText(string(bs))
		input = mod.Entry
		if input == "" {
			input = module2.DEFAULT_ENTRY
		}
	}
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
	ev := eval.New().WithBasePath(input)
	ev.DoSafely(ast).PrintPanic()
}
