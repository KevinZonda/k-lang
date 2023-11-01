package compressor_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/compressor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/jout"
	"testing"
)

func TestCompressAst(t *testing.T) {
	v := parserHelper.Ast("x := 11")
	r := compressor.Compress(v)
	fmt.Println(r)
	fmt.Println("Size: ", len(r))
	ast := compressor.Decompress(r)
	jout.Println(ast)
}
