package visitor

import (
	"errors"
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

type AntlrVisitor struct {
	*parser.BaseV2ParserVisitor
	Errs []VisitorError
}

func New() *AntlrVisitor {
	return &AntlrVisitor{}
}

func (v *AntlrVisitor) appendErr(ctx antlr.ParserRuleContext, msg string, raw error) {
	if raw != nil {
		raw = errors.New(msg)
	}
	e := VisitorError{
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
		EndLine:   ctx.GetStop().GetLine(),
		EndColumn: ctx.GetStop().GetColumn() + len([]rune(ctx.GetStop().GetText())),
		Msg:       msg,
		Text:      ctx.GetText(),
		Raw:       fmt.Errorf(msg),
	}
	v.Errs = append(v.Errs, e)
}

func (v *AntlrVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	var block []node.Node
	for _, t := range ctx.GetChildren() {
		switch t.(type) {
		case *parser.OpenBlockContext:
			ox := v.visitOpenBlock(t.(*parser.OpenBlockContext))
			if ox != nil && len(ox.Openers) > 0 {
				block = append(block, ox)
			}
		case *parser.StructBlockContext:
			if sx := v.visitStructBlock(t.(*parser.StructBlockContext)); sx != nil {
				block = append(block, sx)
			}
		case *parser.FuncBlockContext:
			if fx := v.visitFuncBlock(t.(*parser.FuncBlockContext)); fx != nil {
				block = append(block, fx)
			}
		case *parser.StmtContext:
			if stmt := v.visitStmt(t.(*parser.StmtContext)); stmt != nil {
				block = append(block, stmt)
			}
		case *parser.ExprContext:
			if expr := v.visitExpr(t.(*parser.ExprContext)); expr != nil {
				block = append(block, expr)
			}
		case *antlr.TerminalNodeImpl:
			continue
		}

	}
	return tree.Ast(block)
}
