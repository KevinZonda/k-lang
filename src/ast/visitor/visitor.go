package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
	"github.com/antlr4-go/antlr/v4"
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
		switch t.(type) {
		case *parser.OpenBlockContext:
			continue
		case *parser.StructBlockContext:
			continue
		case *parser.FuncBlockContext:
			continue
		case *parser.StmtContext:
			stmt := v.VisitStmt(t.(*parser.StmtContext)).(node.Stmt)
			block = append(block, stmt)
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
