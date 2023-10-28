package parserHelper

import (
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/parser"
	"github.com/antlr4-go/antlr/v4"
)

type visitor struct {
	*antlr.BaseParseTreeVisitor
}

func newVisitor() *visitor {
	return &visitor{}
}

func (v *visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	return &ast.Ast{}
}
