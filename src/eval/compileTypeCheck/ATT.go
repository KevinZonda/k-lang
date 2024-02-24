package compileTypeCheck

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	. "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/compileTypeCheck/common"
	TMem "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/compileTypeCheck/tmem"
)

type ATT struct {
	m *TMem.Memory
}

func NewATT() *ATT {
	return &ATT{m: TMem.NewMemory()}
}

func (a *ATT) CheckAST(ast tree.Ast) {
	for _, n := range ast {
		switch nT := n.(type) {
		case *node.CodeBlock:
			a.CheckCodeBlock(nT)
		case node.Expr:
			a.CheckExpr(nT)
		case node.Stmt:
			a.CheckStmt(nT)
		case *node.FuncBlock:
			a.CheckFuncBlock(nT)
		case *node.StructBlock:
			a.CheckStructBlock(nT)
		default:
			panic("not implemented")
		}
	}
}

func (a *ATT) CheckFuncBlock(n *node.FuncBlock) {
	a.frameStart(true)
	defer a.frameEnd()
	for _, arg := range n.Args {
		a.m.Top().Set(arg.Name.Value, a.TxType(arg.Type))
	}
	// retT := a.TxType(n.RetType)
	a.CheckCodeBlock(n.Body)
}

func (a *ATT) TxType(t *node.Type) T {
	return T{}
}

func (a *ATT) CheckStructBlock(n *node.StructBlock) {

}

func (a *ATT) CheckCodeBlock(n *node.CodeBlock) {

}

func (a *ATT) CheckStmt(n node.Stmt) {

}

func (a *ATT) CheckAssign(n *node.AssignStmt) {
	rightT := a.CheckExpr(n.Value)
	if len(n.Assignee) == 1 {
		if n.Assignee[0].Type == nil {
			a.m.Top().Set(n.Assignee[0].Var.Value[0].Name.Value, rightT)
			return
		}
		tx := a.TxType(n.Assignee[0].Type)
		if rightT.AssignTo(tx) {
			a.m.Top().Set(n.Assignee[0].Var.Value[0].Name.Value, tx)
			return
		}

	}
}

func (a *ATT) CheckExpr(n node.Expr) T {
	return T{}
}
