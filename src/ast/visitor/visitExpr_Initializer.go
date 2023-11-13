package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitMapInitializer(ctx *parser.MapInitializerContext) interface{} {
	pairs := ctx.AllMapPair()
	var arr []*node.MapPairLiteral
	for _, e := range pairs {
		arr = append(arr, v.VisitMapPairInitializer(e.(*parser.MapPairContext)).(*node.MapPairLiteral))
	}
	return &node.MapLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Value: arr,
	}
}

func (v *AntlrVisitor) VisitMapPairInitializer(ctx *parser.MapPairContext) interface{} {
	if ctx.MapPair() != nil {
		return v.VisitMapPairInitializer(ctx.MapPair().(*parser.MapPairContext))
	}
	return &node.MapPairLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Key:   v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(node.Expr),
		Value: v.VisitExprWithLambda(ctx.ExprWithLambda().(*parser.ExprWithLambdaContext)).(node.Expr),
	}
}

func (v *AntlrVisitor) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	exprs := ctx.AllExpr()
	var arr []node.Expr
	for _, e := range exprs {
		arr = append(arr, v.VisitExpr(e.(*parser.ExprContext)).(node.Expr))
	}
	return &node.ArrayLiteral{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Value: arr,
	}
}
