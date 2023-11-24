package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitStructBlock(ctx parser.IStructBlockContext) *node.StructBlock {
	return &node.StructBlock{
		Token: token.FromAntlrToken(ctx.Struct().GetSymbol()),
		Body:  v.visitDeclareBlock(ctx.DeclareBlock()),
		Name:  v.visitIdentifier(ctx.Identifier()).Value,
	}
}

func (v *AntlrVisitor) visitDeclareBlock(ctx parser.IDeclareBlockContext) map[string]*node.Type {
	if ctx == nil {
		return nil
	}
	var body = make(map[string]*node.Type)
	declares := ctx.AllDeclareStmt()
	for _, declare := range declares {
		tN := v.visitType(declare.Type_())
		if tN != nil {
			for _, id := range declare.AllIdentifier() {
				body[id.GetText()] = tN
			}
		}
	}
	return body
}
