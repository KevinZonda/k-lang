package funcs

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/compressor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
)

func Run(input string) {
	bs, e := iox.ReadAllByte(input)
	panicx.PanicIfNotNil(e, e)

	ast := compressor.Decompress(bs)
	ev := eval.New(ast)
	ev.Do()
	fmt.Println("Evaluated ->", ev.It())
}
