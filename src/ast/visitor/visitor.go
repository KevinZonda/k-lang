package visitor

import (
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

func (v *AntlrVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	var block []node.Node
	for _, t := range ctx.GetChildren() {
		switch t.(type) {
		case *parser.OpenBlockContext:
			ox := v.VisitOpenBlock(t.(*parser.OpenBlockContext)).(*node.OpenBlock)
			if len(ox.Openers) > 0 {
				block = append(block, ox)
			}
		case *parser.StructBlockContext:
			continue
		case *parser.FuncBlockContext:
			fx := v.VisitFuncBlock(t.(*parser.FuncBlockContext)).(node.Block)
			block = append(block, fx)
		case *parser.StmtContext:
			stmt := v.VisitStmt(t.(*parser.StmtContext)).(node.Stmt)
			block = append(block, stmt)
		case *parser.ExprContext:
			expr := v.VisitExpr(t.(*parser.ExprContext)).(node.Expr)
			block = append(block, expr)
		case *antlr.TerminalNodeImpl:
			continue
		}

	}
	return tree.Ast(block)
}
