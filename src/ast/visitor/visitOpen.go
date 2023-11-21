package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitOpenStmt(ctx *parser.OpenStmtContext) interface{} {
	txt := ctx.StringLiteral().GetText()

	o := &node.OpenStmt{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Path:  txt[1 : len(txt)-1],
	}
	if ctx.As() == nil {

		return o
	}
	if as := v.visitIdentifier(ctx.Identifier()); as != nil {
		o.As = as.Value
	}
	return o
}

func (v *AntlrVisitor) VisitOpenBlock(ctx *parser.OpenBlockContext) interface{} {
	os := ctx.AllOpenStmt()

	var openers []*node.OpenStmt
	for _, o := range os {
		openers = append(openers, v.VisitOpenStmt(o.(*parser.OpenStmtContext)).(*node.OpenStmt))
	}
	return &node.OpenBlock{
		Token:   token.FromAntlrToken(ctx.GetStart()),
		Openers: openers,
	}
}
