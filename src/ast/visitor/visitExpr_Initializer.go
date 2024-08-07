package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func (v *AntlrVisitor) visitMapInitializer(ctx parser.IMapInitializerContext) *node.MapLiteral {
	pairs := ctx.AllMapPair()
	var arr []*node.MapPairLiteral
	for _, e := range pairs {
		arr = append(arr, v.visitMapPairInitializer(e))
	}
	return &node.MapLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Value: arr,
	}
}

func (v *AntlrVisitor) visitMapPairInitializer(ctx parser.IMapPairContext) *node.MapPairLiteral {
	if ctx.MapPair() != nil {
		return v.visitMapPairInitializer(ctx.MapPair())
	}
	if ctx.IdentifierPair() != nil {
		k, v, _ := v.visitIdentifierPair(ctx.IdentifierPair())
		return &node.MapPairLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Key:   &node.Identifier{Token: token.FromAntlrToken(ctx.GetStart()), Value: k},
			Value: v,
		}
	}
	return &node.MapPairLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Key:   v.visitExpr(ctx.Expr()),
		Value: v.visitExprWithLambda(ctx.ExprWithLambda()),
	}
}

func (v *AntlrVisitor) visitArrayInitializer(ctx parser.IArrayInitializerContext) *node.ArrayLiteral {
	exprs := ctx.AllExpr()
	var arr []node.Expr
	for _, e := range exprs {
		arr = append(arr, v.visitExpr(e))
	}
	return &node.ArrayLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Value: arr,
	}
}

func (v *AntlrVisitor) visitStructInitializer(ctx parser.IStructInitializerContext) *node.StructLiteral {
	pairs := ctx.AllIdentifierPair()
	body := orderedmap.New[string, node.Expr]()
	for _, e := range pairs {
		if k, val, ok := v.visitIdentifierPair(e); ok {
			body.Set(k, val)
		}
	}
	return &node.StructLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Body:  body,
		Type:  v.visitType(ctx.Type_()),
	}
}

func (v *AntlrVisitor) visitIdentifierPair(ctx parser.IIdentifierPairContext) (key string, val node.Expr, ok bool) {
	if ctx == nil {
		return "", nil, false
	}
	return v.visitIdentifier(ctx.Identifier()).Value, v.visitExprWithLambda(ctx.ExprWithLambda()), true

}
