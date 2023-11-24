package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitExprWithLambda(ctx *parser.ExprWithLambdaContext) node.Expr {
	if ctx.Lambda() != nil {
		return v.visitLambda(ctx.Lambda().(*parser.LambdaContext))
	}
	return v.visitExpr(ctx.Expr().(*parser.ExprContext))
}

func (v *AntlrVisitor) visitDotExpr(ctx *parser.ExprContext) *node.DotExpr {
	return &node.DotExpr{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Left:  v.visitExpr(ctx.GetLHS().(*parser.ExprContext)),
		Right: v.visitExpr(ctx.GetRHS().(*parser.ExprContext)),
	}
}

func (v *AntlrVisitor) visitExpr(ctx *parser.ExprContext) node.Expr {
	if ctx == nil {
		return nil
	}
	// fmt.Println("VisitExpr")
	if ctx.AssignStmt() != nil {
		return v.visitAssignStmt(ctx.AssignStmt().(*parser.AssignStmtContext))
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
		return v.visitFuncCall(ctx.FuncCall().(*parser.FuncCallContext))
	}
	if ctx.Dot() != nil {
		return v.visitDotExpr(ctx)
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

func (v *AntlrVisitor) visitUnaryExpr(ctx *parser.ExprContext) *node.UnaryOperExpr {

	if ctx.GetChildCount() != 2 {
		v.Errs = append(v.Errs, VisitorError{
			Raw:       nil,
			Msg:       "Expected unary expression (e.g., -a), but got " + ctx.GetStart().GetText() + "\nWe expect 2 children, but got " + fmt.Sprint(ctx.GetChildCount()) + " children",
			Line:      ctx.GetStart().GetLine(),
			Column:    ctx.GetStart().GetColumn(),
			EndLine:   ctx.GetStop().GetLine(),
			EndColumn: ctx.GetStop().GetColumn(),
			Text:      ctx.GetText(),
		})
		return nil
	}
	oper := ctx.GetChild(0).(*parser.UnaryOperContext)

	return &node.UnaryOperExpr{
		Token: token.FromAntlrToken(oper.GetStart()),
		Oper:  oper.GetText(),
		Expr:  v.visitExpr(ctx.GetChild(1).(*parser.ExprContext)),
	}
}

func (v *AntlrVisitor) visitBinaryExpr(ctx *parser.ExprContext) *node.BinaryOperExpr {
	// fmt.Println("VisitBinaryExpr")
	if ctx.GetChildCount() != 3 {
		v.Errs = append(v.Errs, VisitorError{
			Raw:       nil,
			Msg:       "Expected binary expression (e.g., a + b), but got " + ctx.GetStart().GetText() + "\nWe expect 3 children, but got " + fmt.Sprint(ctx.GetChildCount()) + " children",
			Line:      ctx.GetStart().GetLine(),
			Column:    ctx.GetStart().GetColumn(),
			EndLine:   ctx.GetStop().GetLine(),
			EndColumn: ctx.GetStop().GetColumn(),
			Text:      ctx.GetText(),
		})
		return nil
	}
	return &node.BinaryOperExpr{
		Token: token.FromAntlrToken(ctx.GetOP()),
		Left:  v.visitExpr(ctx.GetLHS().(*parser.ExprContext)),
		Oper:  ctx.GetOP().GetText(),
		Right: v.visitExpr(ctx.GetRHS().(*parser.ExprContext)),
	}
}
