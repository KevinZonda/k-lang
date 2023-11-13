package compressor_test

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/compressor"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/utils/jout"
	"testing"
)

func TestCompressAst(t *testing.T) {
	v, _ := parserHelper.Ast(`
fn foo(int x) {
  return x + x
}

a := foo(11)

`)
	r, _ := compressor.Compress(v)
	fmt.Println("Compressed To: ", r)
	fmt.Println("Size: ", len(r))
	ast, _ := compressor.Decompress(r)
	jout.Println(ast)
}
