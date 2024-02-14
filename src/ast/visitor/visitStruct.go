package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func (v *AntlrVisitor) visitStructBlock(ctx parser.IStructBlockContext) *node.StructBlock {
	return &node.StructBlock{
		Token: token.FromAntlrToken(ctx.Struct().GetSymbol()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Body:  v.visitDeclareBlock(ctx.DeclareBlock()),
		Name:  v.visitIdentifier(ctx.Identifier()).Value,
	}
}

func (v *AntlrVisitor) visitDeclareBlock(ctx parser.IDeclareBlockContext) *orderedmap.OrderedMap[string, *node.Declare] {
	if ctx == nil {
		return nil
	}
	body := orderedmap.New[string, *node.Declare]()
	declares := ctx.AllDeclareStmt()
	for _, declare := range declares {
		ds := v.visitDeclareStmt(declare)
		for _, name := range ds.Names {
			body.Set(name, &ds.Declare)
		}
	}
	return body
}

func (v *AntlrVisitor) visitDeclareStmt(ctx parser.IDeclareStmtContext) *node.DeclareStmt {
	if ctx.FuncBlock() != nil {
		block := v.visitFuncBlock(ctx.FuncBlock())
		return &node.DeclareStmt{
			Names: []string{block.Name.Value},
			Declare: node.Declare{
				Func: block,
				Type: v.visitType(ctx.Type_()),
			},
		}
	}
	if ctx.ExprWithLambda() != nil {
		return &node.DeclareStmt{
			Names: []string{ctx.Identifier(0).GetText()},
			Declare: node.Declare{
				Type:  v.visitType(ctx.Type_()),
				Value: v.visitExprWithLambda(ctx.ExprWithLambda()),
			},
		}
	}
	stmt := &node.DeclareStmt{
		Declare: node.Declare{
			Type: v.visitType(ctx.Type_()),
		},
	}
	ids := ctx.AllIdentifier()
	for _, id := range ids {
		stmt.Names = append(stmt.Names, id.GetText())
	}
	return stmt
}
