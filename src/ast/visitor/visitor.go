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
			sx := v.VisitStructBlock(t.(*parser.StructBlockContext)).(*node.StructBlock)
			block = append(block, sx)
		case *parser.FuncBlockContext:
			if fx := v.VisitFuncBlock(t.(*parser.FuncBlockContext)); fx != nil {
				block = append(block, fx.(node.Block))
			}
		case *parser.StmtContext:
			if stmt := v.VisitStmt(t.(*parser.StmtContext)); stmt != nil {
				block = append(block, stmt.(node.Stmt))
			}
		case *parser.ExprContext:
			expr := v.VisitExpr(t.(*parser.ExprContext)).(node.Expr)
			block = append(block, expr)
		case *antlr.TerminalNodeImpl:
			continue
		}

	}
	return tree.Ast(block)
}
