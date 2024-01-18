package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"reflect"
	"strconv"
)

func (v *AntlrVisitor) visitLiteral(ctx parser.ILiteralContext) node.Expr {
	if ctx.IntegerLiteral() != nil {
		val, err := strconv.Atoi(ctx.GetStart().GetText())

		if err != nil {
			v.appendErr(ctx, "Expected an integer value, but got "+ctx.GetStart().GetText(), err)
			// FIXME: return nil?
		}
		return &node.IntLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: val,
		}
	}
	if ctx.NumberLiteral() != nil {
		val, err := strconv.ParseFloat(ctx.GetStart().GetText(), 64)
		if err != nil {
			v.appendErr(ctx, "Expected an number value, but got "+ctx.GetStart().GetText(), err)
			return nil
		}
		return &node.FloatLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: val,
		}
	}
	if ctx.StringLiteral() != nil {
		txt := ctx.GetStart().GetText()
		return &node.StringLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: txt[1 : len(txt)-1],
		}
	}
	if ctx.True() != nil {
		return &node.BoolLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: true,
		}
	}
	if ctx.False() != nil {
		return &node.BoolLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: false,
		}
	}

	//switch ctx.GetStart().GetTokenType() {
	//case parser.V2ParserIdentifier:
	//	return &node.Identifier{
	//		Token: token.FromAntlrToken(ctx.GetStart()),
	//		Value: ctx.GetStart().GetText(),
	//	}
	//}
	fmt.Println("VisitLiteralBlock : Unknown type")
	fmt.Println("-> text     : ", ctx.GetStart().GetText())
	fmt.Println("-> tokenType: ", ctx.GetStart().GetTokenType())
	fmt.Println("-> reflect  :", reflect.TypeOf(ctx))
	panic("VisitLiteralBlock : Unknown type")
}

func (v *AntlrVisitor) visitInitializer(ctx parser.IInitializerContext) node.Expr {
	if ctx.ArrayInitializer() != nil {
		return v.visitArrayInitializer(ctx.ArrayInitializer())
	}
	if ctx.MapInitializer() != nil {
		return v.visitMapInitializer(ctx.MapInitializer())
	}

	fmt.Println("-> text     : ", ctx.GetStart().GetText())
	fmt.Println("-> tokenType: ", ctx.GetStart().GetTokenType())
	fmt.Println("-> reflect  :", reflect.TypeOf(ctx))
	panic("VisitInitializer : Unknown type")
}
