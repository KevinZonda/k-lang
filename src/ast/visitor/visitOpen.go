package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitOpenStmt(ctx parser.IOpenStmtContext) *node.OpenStmt {
	txt := ctx.StringLiteral().GetText()

	o := &node.OpenStmt{
		Token: token.FromAntlrToken(ctx.GetStart()).WithEnd(ctx.GetStop()),
		Path:  txt[1 : len(txt)-1],
	}
	if ctx.As() == nil {
		return o
	}
	if ctx.Dot() != nil {
		o.As = "."
	}
	if as := v.visitIdentifier(ctx.Identifier()); as != nil {
		o.As = as.Value
	}
	return o
}
