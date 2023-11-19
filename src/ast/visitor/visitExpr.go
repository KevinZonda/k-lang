package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitExprWithLambda(ctx *parser.ExprWithLambdaContext) interface{} {
	if ctx.Lambda() != nil {
		return v.VisitLambda(ctx.Lambda().(*parser.LambdaContext))
	}
	return v.VisitExpr(ctx.Expr().(*parser.ExprContext))
}

func (v *AntlrVisitor) VisitDotExpr(ctx *parser.ExprContext) interface{} {
	return &node.DotExpr{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Left:  v.VisitExpr(ctx.GetLHS().(*parser.ExprContext)).(node.Expr),
		Right: v.VisitExpr(ctx.GetRHS().(*parser.ExprContext)).(node.Expr),
	}
}

func (v *AntlrVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	if ctx == nil {
		return nil
	}
	// fmt.Println("VisitExpr")
	if ctx.AssignStmt() != nil {
		return v.VisitAssignStmt(ctx.AssignStmt().(*parser.AssignStmtContext))
	}
	if ctx.GetOP() != nil {
		return v.VisitBinaryExpr(ctx)
	}
	if ctx.UnaryOper() != nil {
		return v.VisitUnaryExpr(ctx)
	}
	if ctx.Literal() != nil {
		return v.VisitLiteral(ctx.GetChild(0).(*parser.LiteralContext))
	}
	if ctx.LParen() != nil {
		return v.VisitExpr(ctx.GetChild(1).(*parser.ExprContext))
	}
	if ctx.FuncCall() != nil {
		return v.VisitFuncCall(ctx.FuncCall().(*parser.FuncCallContext))
	}
	if ctx.Dot() != nil {
		return v.VisitDotExpr(ctx)
	}

	if ctx.Identifier() != nil {
		return v.visitIdentifier(ctx.Identifier())
	}

	//for _, tx := range ctx.GetChildren() {
	//	fmt.Println("->", reflect.TypeOf(tx))
	//	switch tx.(type) {
	//	case *parser.ExprContext:
	//		_ = v.VisitExpr(tx.(*parser.ExprContext))
	//		continue
	//	case *parser.LiteralContext:
	//		_ = v.VisitLiteral(tx.(*parser.LiteralContext))
	//		continue
	//	case *parser.LambdaContext:
	//		continue
	//	case *parser.VarContext:
	//		continue
	//	}
	//}
	return nil
}

func (v *AntlrVisitor) VisitUnaryExpr(ctx *parser.ExprContext) interface{} {

	if ctx.GetChildCount() != 2 {
		v.Errs = append(v.Errs, VisitorError{
			Raw:    nil,
			Msg:    "Expected unary expression (e.g., -a), but got " + ctx.GetStart().GetText() + "\nWe expect 2 children, but got " + fmt.Sprint(ctx.GetChildCount()) + " children",
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Text:   ctx.GetText(),
		})
		return nil
	}
	oper := ctx.GetChild(0).(*parser.UnaryOperContext)

	return &node.UnaryOperExpr{
		Token: token.FromAntlrToken(oper.GetStart()),
		Oper:  oper.GetText(),
		Expr:  v.VisitExpr(ctx.GetChild(1).(*parser.ExprContext)).(node.Expr),
	}
}

func (v *AntlrVisitor) VisitBinaryExpr(ctx *parser.ExprContext) interface{} {
	// fmt.Println("VisitBinaryExpr")
	if ctx.GetChildCount() != 3 {
		v.Errs = append(v.Errs, VisitorError{
			Raw:    nil,
			Msg:    "Expected binary expression (e.g., a + b), but got " + ctx.GetStart().GetText() + "\nWe expect 3 children, but got " + fmt.Sprint(ctx.GetChildCount()) + " children",
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Text:   ctx.GetText(),
		})
		return nil
	}
	return &node.BinaryOperExpr{
		Token: token.FromAntlrToken(ctx.GetOP()),
		Left:  v.VisitExpr(ctx.GetLHS().(*parser.ExprContext)).(node.Expr),
		Oper:  ctx.GetOP().GetText(),
		Right: v.VisitExpr(ctx.GetRHS().(*parser.ExprContext)).(node.Expr),
	}
}
