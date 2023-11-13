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

func (v *AntlrVisitor) VisitLambda(ctx *parser.LambdaContext) interface{} {
	fb := &node.LambdaExpr{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Body:  v.VisitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)).(*node.CodeBlock),
	}
	args := v.VisitFuncSignArgs(toPtr[parser.FuncSignArgsContext](ctx.FuncSignArgs()))
	ret := v.VisitType(toPtr[parser.TypeContext](ctx.Type_()))
	if args != nil {
		fb.Args = args.([]*node.FuncArg)
	}
	if ret != nil {
		fb.RetType = ret.(*node.Identifier)
	}
	return fb
}

func (v *AntlrVisitor) VisitFuncSig(ctx *parser.FuncSigContext) interface{} {
	fb := &node.FuncBlock{
		// Token: token.FromAntlrToken(ctx.GetStart()),
		Name: v.visitIdentifier(ctx.Identifier()),
	}

	args := v.VisitFuncSignArgs(toPtr[parser.FuncSignArgsContext](ctx.FuncSignArgs()))
	ret := v.VisitType(toPtr[parser.TypeContext](ctx.Type_()))
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
	ts := ctx.AllFuncSignArgItem()
	var rst []*node.FuncArg
	for _, t := range ts {
		rst = append(rst, v.VisitFuncSignArgItem(t.(*parser.FuncSignArgItemContext)).(*node.FuncArg))
	}
	return rst
}

func (v *AntlrVisitor) VisitFuncSignArgItem(ctx *parser.FuncSignArgItemContext) any {
	a := node.FuncArg{
		Name: v.visitIdentifier(ctx.Identifier()),
	}
	if ctx.Type_() != nil {
		a.Type = v.VisitType(ctx.Type_().(*parser.TypeContext)).(*node.Identifier)
	}
	return &a
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
