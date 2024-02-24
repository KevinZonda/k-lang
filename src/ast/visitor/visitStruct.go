package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
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
			},
		}
	}
	var _type parser.ITypeContext
	var _names []antlr.TerminalNode
	if ctx.TypedIdentifiers() != nil {
		_type = ctx.TypedIdentifiers().Type_()
		_names = ctx.TypedIdentifiers().AllIdentifier()
	} else if ctx.TypedIdentifier() != nil {
		_type = ctx.TypedIdentifier().Type_()
		_names = []antlr.TerminalNode{ctx.TypedIdentifier().Identifier()}
	} else {
		v.appendErr(ctx, "Unknown declare", nil)
		return nil
	}

	if ctx.ExprWithLambda() != nil {
		return &node.DeclareStmt{
			Names: []string{_names[0].GetText()},
			Declare: node.Declare{
				Type:  v.visitType(_type),
				Value: v.visitExprWithLambda(ctx.ExprWithLambda()),
			},
		}
	}
	stmt := &node.DeclareStmt{
		Declare: node.Declare{
			Type: v.visitType(_type),
		},
	}
	for _, id := range _names {
		stmt.Names = append(stmt.Names, id.GetText())
	}
	return stmt
}
