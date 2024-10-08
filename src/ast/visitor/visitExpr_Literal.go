package visitor

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"reflect"
	"strconv"
	"strings"
)

func (v *AntlrVisitor) visitLiteral(ctx parser.ILiteralContext) node.Expr {
	if ctx.Nil() != nil {
		return &node.NilLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		}
	}
	if ctx.IntegerLiteral() != nil {
		val, err := strconv.Atoi(ctx.GetStart().GetText())

		if err != nil {
			v.appendErr(ctx, "Expected an integer value, but got "+ctx.GetStart().GetText(), err)
			return nil
		}
		return &node.IntLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
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
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Value: val,
		}
	}
	if ctx.StringLiteral() != nil {
		txt := ctx.GetStart().GetText()
		modeRune := txt[0]
		str := node.StringLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		}
		switch modeRune {
		case '@', '$':
			str.Mode = modeRune
			str.Value = txt[2 : len(txt)-1]
			str.Char = txt[1]
		default:
			str.Mode = ' '
			str.Value = txt[1 : len(txt)-1]
			str.Char = txt[0]
		}
		return &str
	}
	if ctx.True() != nil {
		return &node.BoolLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
			Value: true,
		}
	}
	if ctx.False() != nil {
		return &node.BoolLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
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

	sb := strings.Builder{}
	sb.WriteString("VisitLiteralBlock : Unknown type\n")
	sb.WriteString(fmt.Sprintln("-> reflect  :", reflect.TypeOf(ctx)))
	v.appendErr(ctx, sb.String(), nil)
	return nil
}

func (v *AntlrVisitor) visitInitializer(ctx parser.IInitializerContext) node.Expr {
	if ctx.ArrayInitializer() != nil {
		return v.visitArrayInitializer(ctx.ArrayInitializer())
	}
	if ctx.MapInitializer() != nil {
		return v.visitMapInitializer(ctx.MapInitializer())
	}

	sb := strings.Builder{}
	sb.WriteString("VisitInitializer : Unknown type\n")
	sb.WriteString(fmt.Sprintln("-> reflect  :", reflect.TypeOf(ctx)))
	v.appendErr(ctx, sb.String(), nil)
	return nil
}
