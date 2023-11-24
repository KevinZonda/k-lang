package visitor

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parser"
)

func (v *AntlrVisitor) visitFuncBlock(ctx *parser.FuncBlockContext) *node.FuncBlock {
	fx := v.visitFuncSig(ctx.FuncSig().(*parser.FuncSigContext))
	fx.Token = token.FromAntlrToken(ctx.GetStart())
	fx.Body = v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext))
	return fx
}

func (v *AntlrVisitor) visitLambda(ctx *parser.LambdaContext) *node.LambdaExpr {
	fb := &node.LambdaExpr{
		Token: token.FromAntlrToken(ctx.GetStart()),
		Body:  v.visitCodeBlock(ctx.CodeBlock().(*parser.CodeBlockContext)),
	}
	fb.Args = v.visitFuncSignArgs(toPtr[parser.FuncSignArgsContext](ctx.FuncSignArgs()))
	fb.RetType = v.visitType(toPtr[parser.TypeContext](ctx.Type_()))
	return fb
}

func (v *AntlrVisitor) visitFuncSig(ctx *parser.FuncSigContext) *node.FuncBlock {
	fb := &node.FuncBlock{
		// Token: token.FromAntlrToken(ctx.GetStart()),
		Name: v.visitIdentifier(ctx.Identifier()),
	}

	fb.Args = v.visitFuncSignArgs(toPtr[parser.FuncSignArgsContext](ctx.FuncSignArgs()))
	fb.RetType = v.visitType(toPtr[parser.TypeContext](ctx.Type_()))
	return fb
}

func (v *AntlrVisitor) visitFuncSignArgs(ctx *parser.FuncSignArgsContext) []*node.FuncArg {
	// []*node.FuncArg
	if ctx == nil {
		return nil
	}
	ts := ctx.AllFuncSignArgItem()
	var rst []*node.FuncArg
	for _, t := range ts {
		rst = append(rst, v.visitFuncSignArgItem(t.(*parser.FuncSignArgItemContext)))
	}
	return rst
}

func (v *AntlrVisitor) visitFuncSignArgItem(ctx *parser.FuncSignArgItemContext) *node.FuncArg {
	a := node.FuncArg{
		Name: v.visitIdentifier(ctx.Identifier()),
	}
	if ctx.Type_() != nil {
		a.Type = v.visitType(ctx.Type_().(*parser.TypeContext))
	}
	return &a
}

func (v *AntlrVisitor) visitCodeBlock(ctx *parser.CodeBlockContext) *node.CodeBlock {
	var body []node.Node
	for _, m := range ctx.GetChildren() {
		switch m.(type) {
		case *parser.StmtContext:
			body = append(body, v.visitStmt(m.(*parser.StmtContext)))
		case *parser.ExprContext:
			body = append(body, v.visitExpr(m.(*parser.ExprContext)))
		}
	}
	return &node.CodeBlock{
		Nodes: body,
		Token: token.FromAntlrToken(ctx.GetStart()),
	}
}
