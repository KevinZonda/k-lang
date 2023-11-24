package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitStructBlock(ctx *parser.StructBlockContext) *node.StructBlock {
	return &node.StructBlock{
		Token: token.FromAntlrToken(ctx.Struct().GetSymbol()),
		Body:  v.visitDeclareBlock(ctx.DeclareBlock().(*parser.DeclareBlockContext)),
		Name:  v.visitIdentifier(ctx.Identifier()).Value,
	}
}

func (v *AntlrVisitor) visitDeclareBlock(ctx *parser.DeclareBlockContext) map[string]*node.Type {
	if ctx == nil {
		return nil
	}
	var body = make(map[string]*node.Type)
	declares := ctx.AllDeclareStmt()
	for _, declare := range declares {
		declareStmt := declare.(*parser.DeclareStmtContext)
		tN := v.visitType(toPtr[parser.TypeContext](declareStmt.Type_()))
		if tN != nil {
			for _, id := range declareStmt.AllIdentifier() {
				body[id.GetText()] = tN
			}
		}
	}
	return body
}
