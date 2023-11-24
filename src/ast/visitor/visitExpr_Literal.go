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
	switch ctx.GetStart().GetTokenType() {
	case parser.V2ParserIdentifier:
		return &node.Identifier{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: ctx.GetStart().GetText(),
		}
	case parser.V2ParserIntegerLiteral:
		val, err := strconv.Atoi(ctx.GetStart().GetText())

		if err != nil {
			v.Errs = append(v.Errs, VisitorError{
				Raw:       err,
				Msg:       "Expected an integer value, but got " + ctx.GetStart().GetText(),
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn(),
				Text:      ctx.GetText(),
			})
		}
		return &node.IntLiteral{
			Token: token.FromAntlrToken(ctx.GetStart()),
			Value: val,
		}
	case parser.V2ParserNumberLiteral:
		val, err := strconv.ParseFloat(ctx.GetStart().GetText(), 64)
		if err != nil {
			v.Errs = append(v.Errs, VisitorError{
				Raw:       err,
				Msg:       "Expected an number value, but got " + ctx.GetStart().GetText(),
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn(),
				Text:      ctx.GetText(),
			})
			return nil
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
	case parser.V2ParserStruct:
		fmt.Println("-> STRUCT : ", ctx.GetStart().GetText())
	default:
		if ctx.ArrayInitializer() != nil {
			return v.visitArrayInitializer(ctx.ArrayInitializer())
		}
		if ctx.MapInitializer() != nil {
			return v.visitMapInitializer(ctx.MapInitializer())
		}
	}
	fmt.Println("VisitLiteralBlock : Unknown type")
	fmt.Println("-> text     : ", ctx.GetStart().GetText())
	fmt.Println("-> tokenType: ", ctx.GetStart().GetTokenType())
	fmt.Println("-> reflect  :", reflect.TypeOf(ctx))
	panic("VisitLiteralBlock : Unknown type")
}
