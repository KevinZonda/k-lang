package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/node"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/token"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/ast/tree"
	"git.cs.bham.ac.uk/xxs166/uob-project/klang/parser"
	"github.com/antlr4-go/antlr/v4"
	"reflect"
	"strconv"
)

type AntlrVisitor struct {
	*parser.BaseV2ParserVisitor
}

func New() *AntlrVisitor {
	return &AntlrVisitor{}
}

func (v *AntlrVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	var block []node.Node
	for _, t := range ctx.GetChildren() {
		// fmt.Println(reflect.TypeOf(t))
		switch t.(type) {
		case *parser.OpenBlockContext:
			continue
		case *parser.StructBlockContext:
			continue
		case *parser.FuncBlockContext:
			continue
		case *parser.StmtContext:
			continue
		case *parser.ExprContext:
			expr := v.VisitExpr(t.(*parser.ExprContext)).(node.Expr)
			block = append(block, expr)
			continue
		case *antlr.TerminalNodeImpl:
			continue
		}

	}
	return tree.Ast(block)
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
	case parser.V2ParserRULE_arrayInitializer:
		fmt.Println("-> ARRAY  : ", ctx.GetStart().GetText())
	case parser.V2ParserRULE_mapInitializer:
		fmt.Println("-> MAP    : ", ctx.GetStart().GetText())
	case parser.V2ParserRULE_structInitializer:
		fmt.Println("-> STRUCT : ", ctx.GetStart().GetText())
	default:
		fmt.Println("VisitLiteralBlock : Unknown type")
		fmt.Println("-> text     : ", ctx.GetStart().GetText())
		fmt.Println("-> tokenType: ", ctx.GetStart().GetTokenType())
		fmt.Println("-> reflect  :", reflect.TypeOf(ctx))
		panic("VisitLiteralBlock : Unknown type")
	}
	panic("Unimplemented")
}
