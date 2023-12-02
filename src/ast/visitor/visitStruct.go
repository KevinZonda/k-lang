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

func (v *AntlrVisitor) visitDeclareBlock(ctx parser.IDeclareBlockContext) map[string]*node.Declare {
	if ctx == nil {
		return nil
	}
	var body = make(map[string]*node.Declare)
	declares := ctx.AllDeclareStmt()
	for _, declare := range declares {
		if declare.ExprWithLambda() != nil {
			body[declare.Identifier(0).GetText()] = &node.Declare{
				Type:  v.visitType(declare.Type_()),
				Value: v.visitExprWithLambda(declare.ExprWithLambda()),
			}
			continue
		}
		tN := v.visitType(declare.Type_())
		if tN != nil {
			for _, id := range declare.AllIdentifier() {
				body[id.GetText()] = &node.Declare{
					Type: tN,
				}
			}
		}
	}
	return body
}
