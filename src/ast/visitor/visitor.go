package visitor

import (
	"errors"
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
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
		Msg: msg,
		Raw: fmt.Errorf(msg),
	}
	if ctx == nil {
		goto end
	}

	e.Text = ctx.GetText()
	if ctx.GetStart() != nil {
		e.Line = ctx.GetStart().GetLine()
		e.Column = ctx.GetStart().GetColumn()
	}
	if ctx.GetStop() != nil {
		e.EndLine = ctx.GetStop().GetLine()
		e.EndColumn = ctx.GetStop().GetColumn() + len([]rune(ctx.GetStop().GetText()))
	}

end:
	v.Errs = append(v.Errs, e)
}

func (v *AntlrVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	var block []node.Node
	for _, t := range ctx.GetChildren() {
		var n node.Node
		switch t.(type) {
		case *parser.StructBlockContext:
			n = v.visitStructBlock(t.(*parser.StructBlockContext))
		case *parser.FuncBlockContext:
			n = v.visitFuncBlock(t.(*parser.FuncBlockContext))
		case *parser.StmtContext:
			n = v.visitStmt(t.(*parser.StmtContext))
		case *parser.ExprContext:
			n = v.visitExpr(t.(*parser.ExprContext))
		case *parser.CommaExprContext:
			n = v.visitCommaExpr(t.(*parser.CommaExprContext))
		}
		if n != nil {
			block = append(block, n)
		}

	}
	return tree.Ast(block)
}

func (v *AntlrVisitor) visitCommaExpr(ctx parser.ICommaExprContext) node.Expr {
	if ctx == nil {
		return nil
	}
	if ctx.GetChildCount() == 0 {
		return nil
	}

	exprs := ctx.AllExprWithLambda()
	commas := ctx.AllComma()
	if len(exprs) == 1 && len(commas) == 0 {
		return v.visitExprWithLambda(exprs[0])
	}
	exprList := make([]node.Expr, 0, len(exprs))
	for _, expr := range exprs {
		exprList = append(exprList, v.visitExprWithLambda(expr))
	}
	return &node.CommaExpr{
		Token: token.FromAntlrToken(ctx.GetStart()).WithBegin(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Exprs: exprList,
	}
}
