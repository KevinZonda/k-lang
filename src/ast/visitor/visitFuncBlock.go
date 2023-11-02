package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) VisitFuncBlock(ctx *parser.FuncBlockContext) interface{} {
	fx := v.VisitFuncSig(ctx.FuncSig().(*parser.FuncSigContext)).(*node.FuncBlock)
	fx.Token = token.FromAntlrToken(ctx.GetStart())
	fx.Body = v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock)
	return fx
}

func (v *AntlrVisitor) VisitFuncSig(ctx *parser.FuncSigContext) interface{} {
	fb := &node.FuncBlock{
		// Token: token.FromAntlrToken(ctx.GetStart()),
		Name: v.visitIdentifier(ctx.Identifier()),
	}

	args := v.VisitFuncSignArgs(typeCastToPtr[parser.FuncSignArgsContext](ctx.FuncSignArgs()))
	ret := v.VisitType(typeCastToPtr[parser.TypeContext](ctx.Type_()))
	if args != nil {
		fb.Args = args.([]*node.FuncArg)
	}
	if ret != nil {
		fb.RetType = ret.(*node.Identifier)
	}
	return fb
}

func (v *AntlrVisitor) VisitFuncSignArgs(ctx *parser.FuncSignArgsContext) interface{} {
	// []*node.FuncArg
	if ctx == nil {
		return nil
	}
	ts := ctx.AllType_()
	vs := ctx.AllIdentifier()
	var rst []*node.FuncArg
	for i, t := range ts {
		rst = append(rst, &node.FuncArg{
			Type: v.VisitType(t.(*parser.TypeContext)).(*node.Identifier),
			Name: v.visitIdentifier(vs[i]),
		})
	}
	return rst
}

func (v *AntlrVisitor) VisitCodeBlock(ctx *parser.CodeBlockContext) interface{} {
	var body []node.Node
	for _, m := range ctx.GetChildren() {
		switch m.(type) {
		case *parser.StmtContext:
			body = append(body, v.VisitStmt(m.(*parser.StmtContext)).(node.Node))
		case *parser.ExprContext:
			body = append(body, v.VisitExpr(m.(*parser.ExprContext)).(node.Node))
		}
	}
	return &node.CodeBlock{
		Nodes: body,
		Token: token.FromAntlrToken(ctx.GetStart()),
	}
}
