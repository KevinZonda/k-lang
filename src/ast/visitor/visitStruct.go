package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitStructBlock(ctx *parser.StructBlockContext) any {
	return &node.StructBlock{
		Token: token.FromAntlrToken(ctx.Struct().GetSymbol()),
		Body:  v.visitDeclareBlock(ctx.DeclareBlock().(*parser.DeclareBlockContext)),
		Name:  v.visitIdentifier(ctx.Identifier()).Value,
	}
}

func (v *AntlrVisitor) visitDeclareBlock(ctx *parser.DeclareBlockContext) map[string]string {
	if ctx == nil {
		return nil
	}
	var body = make(map[string]string)
	declares := ctx.AllDeclareStmt()
	for _, declare := range declares {
		declareStmt := declare.(*parser.DeclareStmtContext)
		tN := v.VisitType(toPtr[parser.TypeContext](declareStmt.Type_()))
		if tN != nil {
			t := tN.(*node.Identifier)
			for _, id := range declareStmt.AllIdentifier() {
				body[id.GetText()] = t.Value
			}
		}
	}
	return body
}
