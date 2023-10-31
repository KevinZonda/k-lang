package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
	"reflect"
	"strconv"
)

func (v *AntlrVisitor) VisitExprWithLambda(ctx *parser.ExprWithLambdaContext) interface{} {
	fmt.Println("VisitExprWithLambda")
	if ctx.Lambda() != nil {
		panic("VisitExprWithLambda : Lambda not implemented")
	}
	return v.VisitExpr(ctx.Expr().(*parser.ExprContext))
}

func (v *AntlrVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	// fmt.Println("VisitExpr")
	if ctx.BinaryOper() != nil {
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
	if ctx.GetChildCount() == 1 {
		// TODO: ID
		n := ctx.GetChild(0).(*antlr.TerminalNodeImpl)
		return &node.Identifier{
			Token: token.FromAntlrToken(n.GetSymbol()),
			Value: n.GetText(),
		}
	}
	//for _, t := range ctx.GetChildren() {
	//	fmt.Println("->", reflect.TypeOf(t))
	//	switch t.(type) {
	//	case *parser.ExprContext:
	//		_ = v.VisitExpr(t.(*parser.ExprContext))
	//		continue
	//	case *parser.LiteralContext:
	//		_ = v.VisitLiteral(t.(*parser.LiteralContext))
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
		panic("VisitUnaryExpr : Invalid child count : " + fmt.Sprint(ctx.GetChildCount()))
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
		panic("VisitBinaryExpr : Invalid child count : " + fmt.Sprint(ctx.GetChildCount()))
	}
	oper := ctx.GetChild(1).(*parser.BinaryOperContext)

	return &node.BinaryOperExpr{
		Token: token.FromAntlrToken(oper.GetStart()),
		Left:  v.VisitExpr(ctx.GetChild(0).(*parser.ExprContext)).(node.Expr),
		Oper:  oper.GetText(),
		Right: v.VisitExpr(ctx.GetChild(2).(*parser.ExprContext)).(node.Expr),
	}
}

func (v *AntlrVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	switch ctx.GetStart().GetTokenType() {
	case parser.V2ParserIdentifier:
		return &node.Identifier{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: ctx.GetStart().GetText(),
		}
	case parser.V2ParserIntegerLiteral:
		val, err := strconv.Atoi(ctx.GetStart().GetText())
		if err != nil {
			panic(err)
		}
		return &node.IntLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: val,
		}
	case parser.V2ParserNumberLiteral:
		val, err := strconv.ParseFloat(ctx.GetStart().GetText(), 64)
		if err != nil {
			panic(err)
		}
		return &node.FloatLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: val,
		}
	case parser.V2ParserStringLiteral:
		txt := ctx.GetStart().GetText()
		return &node.StringLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: txt[1 : len(txt)-1],
		}
	case parser.V2ParserTrue:
		return &node.BoolLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: true,
		}
	case parser.V2ParserFalse:
		return &node.BoolLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: false,
		}
	case parser.V2ParserRULE_structInitializer:
		fmt.Println("-> STRUCT : ", ctx.GetStart().GetText())
	default:
		if ctx.ArrayInitializer() != nil {
			return v.VisitArrayInitializer(ctx.ArrayInitializer().(*parser.ArrayInitializerContext))
		}
		if ctx.MapInitializer() != nil {
			return v.VisitMapInitializer(ctx.MapInitializer().(*parser.MapInitializerContext))
		}
		fmt.Println("VisitLiteralBlock : Unknown type")
		fmt.Println("-> text     : ", ctx.GetStart().GetText())
		fmt.Println("-> tokenType: ", ctx.GetStart().GetTokenType())
		fmt.Println("-> reflect  :", reflect.TypeOf(ctx))
		panic("VisitLiteralBlock : Unknown type")
	}
	panic("Unimplemented")
}

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