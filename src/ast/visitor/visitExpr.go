package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitExprWithLambda(ctx parser.IExprWithLambdaContext) node.Expr {
	if ctx == nil {
		return nil
	}
	if ctx.GetTriCond() != nil {
		return &node.TrinaryOperExpr{
			Token:   token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Cond:    v.visitExpr(ctx.GetTriCond()),
			IfTrue:  v.visitExprWithLambda(ctx.GetIfTrue()),
			IfFalse: v.visitExprWithLambda(ctx.GetIfFalse()),
		}
	}
	if ctx.GetParenExpr() != nil {
		return v.visitExprWithLambda(ctx.GetParenExpr())
	}
	if ctx.GetCallExpr() != nil {
		return &node.FuncCall{
			Token:    token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			CallExpr: v.visitExprWithLambda(ctx.GetCallExpr()),
			Args:     v.visitFuncCallArgs(ctx.FuncCallArgs()),
		}
	}
	if ctx.Lambda() != nil {
		return v.visitLambda(ctx.Lambda())
	}
	return v.visitExpr(ctx.Expr())
}

func (v *AntlrVisitor) visitRefExpr(ctx parser.IExprContext) *node.RefExpr {
	return &node.RefExpr{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Expr:  v.visitExpr(ctx.Expr(0)),
	}
}

func (v *AntlrVisitor) visitDotExpr(ctx parser.IExprContext) *node.DotExpr {
	if ctx == nil {
		return nil
	}
	return &node.DotExpr{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Left:  v.visitExpr(ctx.GetLHS()),
		Right: v.visitExpr(ctx.GetRHS()),
	}
}

func (v *AntlrVisitor) visitIndexExpr(ctx parser.IExprContext) *node.IndexExpr {
	return &node.IndexExpr{
		Token:    token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Left:     v.visitExpr(ctx.GetLHS()),
		Index:    v.visitExpr(ctx.GetIndex()),
		EndIndex: v.visitExpr(ctx.GetEndIndex()),
		Col:      ctx.Col() != nil,
	}
}

func (v *AntlrVisitor) visitExprs(ctxs []parser.IExprContext) []node.Expr {
	var exprs []node.Expr
	for _, ctx := range ctxs {
		e := v.visitExpr(ctx)
		if e != nil {
			exprs = append(exprs, e)
		}
	}
	return exprs

}

func (v *AntlrVisitor) visitExpr(ctx parser.IExprContext) node.Expr {
	if ctx == nil {
		return nil
	}
	// fmt.Println("VisitExpr")
	if ctx.GetCallExpr() != nil {
		return &node.FuncCall{
			Token:    token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			CallExpr: v.visitExpr(ctx.GetCallExpr()),
			Args:     v.visitFuncCallArgs(ctx.FuncCallArgs()),
		}
	}
	if ctx.Ref() != nil {
		return v.visitRefExpr(ctx)
	}
	if ctx.GetTriCond() != nil {
		return &node.TrinaryOperExpr{
			Token:   token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Cond:    v.visitExpr(ctx.GetTriCond()),
			IfTrue:  v.visitExpr(ctx.GetIfTrue()),
			IfFalse: v.visitExpr(ctx.GetIfFalse()),
		}

	}
	if ctx.GetOP() != nil {
		return v.visitBinaryExpr(ctx)
	}
	if ctx.UnaryOper() != nil {
		return v.visitUnaryExpr(ctx)
	}
	if ctx.Literal() != nil {
		return v.visitLiteral(ctx.Literal())
	}
	if ctx.LParen() != nil {
		return v.visitExpr(ctx.GetChild(1).(*parser.ExprContext))
	}
	if ctx.FuncCall() != nil {
		return v.visitFuncCall(ctx.FuncCall())
	}
	if ctx.Dot() != nil {
		return v.visitDotExpr(ctx)
	}
	if ctx.GetIndex() != nil {
		return v.visitIndexExpr(ctx)
	}
	if ctx.Initializer() != nil {
		return v.visitInitializer(ctx.Initializer())
	}
	if ctx.StructInitializer() != nil {
		return v.visitStructInitializer(ctx.StructInitializer())
	}

	if ctx.Identifier() != nil {
		return v.visitIdentifier(ctx.Identifier())
	}

	//for _, tx := range ctx.GetChildren() {
	//	fmt.Println("->", reflect.TypeOf(tx))
	//	switch tx.(type) {
	//	case *parser.ExprContext:
	//		_ = v.visitExpr(tx.(*parser.ExprContext))
	//		continue
	//	case *parser.LiteralContext:
	//		_ = v.visitLiteral(tx.(*parser.LiteralContext))
	//		continue
	//	case *parser.LambdaContext:
	//		continue
	//	case *parser.VarContext:
	//		continue
	//	}
	//}
	return nil
}

func (v *AntlrVisitor) visitUnaryExpr(ctx parser.IExprContext) *node.UnaryOperExpr {
	if ctx == nil {
		return nil
	}
	if ctx.GetChildCount() != 2 {
		v.appendErr(ctx,
			"Expected unary expression (e.g., -a), but got "+ctx.GetStart().GetText()+"\nWe expect 2 children, but got "+fmt.Sprint(ctx.GetChildCount())+" children",
			nil)

		return nil
	}
	oper := ctx.UnaryOper()

	return &node.UnaryOperExpr{
		Token: token.FromAntlrToken(oper.GetStart()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Oper:  oper.GetText(),
		Expr:  v.visitExpr(ctx.GetChild(1).(parser.IExprContext)),
	}
}

func (v *AntlrVisitor) visitBinaryExpr(ctx parser.IExprContext) *node.BinaryOperExpr {
	// fmt.Println("VisitBinaryExpr")
	if ctx.GetChildCount() != 3 {
		v.appendErr(ctx,
			"Expected binary expression (e.g., a + b), but got "+ctx.GetStart().GetText()+"\nWe expect 3 children, but got "+fmt.Sprint(ctx.GetChildCount())+" children",
			nil)
		return nil
	}
	return &node.BinaryOperExpr{
		Token: token.FromAntlrToken(ctx.GetOP()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Left:  v.visitExpr(ctx.GetLHS()),
		Oper:  ctx.GetOP().GetText(),
		Right: v.visitExpr(ctx.GetRHS()),
	}
}
