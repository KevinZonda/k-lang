package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/klang/parser"
)

func (v *AntlrVisitor) VisitExprWithLambda(ctx *parser.ExprWithLambdaContext) interface{} {
	fmt.Println("VisitExprWithLambda")
	if ctx.Lambda() != nil {
		panic("VisitExprWithLambda : Lambda not implemented")
	}
	return v.VisitExpr(ctx.Expr().(*parser.ExprContext))
}
