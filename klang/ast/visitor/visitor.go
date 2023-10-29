package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/tree"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/parser"
	"reflect"
)

type AntlrVisitor struct {
	*parser.BaseV2ParserVisitor
}

func New() *AntlrVisitor {
	return &AntlrVisitor{}
}

func (v *AntlrVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	fmt.Println("VisitProgram")
	fmt.Println(ctx.GetText())
	fmt.Println(reflect.TypeOf(ctx))
	return &tree.Ast{}
}
