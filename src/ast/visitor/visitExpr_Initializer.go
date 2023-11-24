package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitMapInitializer(ctx *parser.MapInitializerContext) *node.MapLiteral {
	pairs := ctx.AllMapPair()
	var arr []*node.MapPairLiteral
	for _, e := range pairs {
		arr = append(arr, v.visitMapPairInitializer(e.(*parser.MapPairContext)))
	}
	return &node.MapLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Value: arr,
	}
}

func (v *AntlrVisitor) visitMapPairInitializer(ctx *parser.MapPairContext) *node.MapPairLiteral {
	if ctx.MapPair() != nil {
		return v.visitMapPairInitializer(ctx.MapPair().(*parser.MapPairContext))
	}
	return &node.MapPairLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Key:   v.visitExpr(ctx.Expr().(*parser.ExprContext)),
		Value: v.visitExprWithLambda(ctx.ExprWithLambda().(*parser.ExprWithLambdaContext)),
	}
}

func (v *AntlrVisitor) visitArrayInitializer(ctx *parser.ArrayInitializerContext) *node.ArrayLiteral {
	exprs := ctx.AllExpr()
	var arr []node.Expr
	for _, e := range exprs {
		arr = append(arr, v.visitExpr(e.(*parser.ExprContext)))
	}
	return &node.ArrayLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Value: arr,
	}
}
