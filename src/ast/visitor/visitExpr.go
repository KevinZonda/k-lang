package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitExprWithLambda(ctx parser.IExprWithLambdaContext) node.Expr {
	if ctx.Lambda() != nil {
		return v.visitLambda(ctx.Lambda())
	}
	return v.visitExpr(ctx.Expr())
}

func (v *AntlrVisitor) visitDotExpr(ctx parser.IExprContext) *node.DotExpr {
	return &node.DotExpr{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Left:  v.visitExpr(ctx.GetLHS()),
		Right: v.visitExpr(ctx.GetRHS()),
	}
}

func (v *AntlrVisitor) visitIndexExpr(ctx parser.IExprContext) *node.IndexExpr {
	return &node.IndexExpr{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Left:  v.visitExpr(ctx.GetLHS()),
		Index: v.visitExpr(ctx.GetIndex()),
	}
}

func (v *AntlrVisitor) visitExpr(ctx parser.IExprContext) node.Expr {
	if ctx == nil {
		return nil
	}
	// fmt.Println("VisitExpr")
	if ctx.AssignStmt() != nil {
		return v.visitAssignStmt(ctx.AssignStmt())
	}
	if ctx.GetOP() != nil {
		return v.visitBinaryExpr(ctx)
	}
	if ctx.UnaryOper() != nil {
		return v.visitUnaryExpr(ctx)
	}
	if ctx.Literal() != nil {
		return v.visitLiteral(ctx.GetChild(0).(*parser.LiteralContext))
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

	if ctx.GetChildCount() != 2 {
		v.appendErr(ctx,
			"Expected unary expression (e.g., -a), but got "+ctx.GetStart().GetText()+"\nWe expect 2 children, but got "+fmt.Sprint(ctx.GetChildCount())+" children",
			nil)

		return nil
	}
	oper := ctx.UnaryOper()

	return &node.UnaryOperExpr{
		Token: token.FromAntlrToken(oper.GetStart()),
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
		Token: token.FromAntlrToken(ctx.GetOP()),
		Left:  v.visitExpr(ctx.GetLHS()),
		Oper:  ctx.GetOP().GetText(),
		Right: v.visitExpr(ctx.GetRHS()),
	}
}
