package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
)

func (e *Eval) EvalWhileForStmt(n *node.WhileStyleFor) any {
	e.loopLvl++
	for {
		if n.ConditionExpr != nil {
			rst := e.EvalExpr(n.ConditionExpr).(bool)
			if !rst {
				e.loopLvl--
				return nil
			}
		}
		_ = e.EvalLoopCodeBlock(n.Body)
		if e.objTable.HasKeyAtTop("0") {
			e.loopLvl--
			return nil
		}
		if e.objTable.HasKeyAtTop("1") {
			e.loopLvl--
			e.objTable.RemoveKeyAtTop("1")
			return nil
		}
	}
}

func (e *Eval) EvalLoopCodeBlock(fc *node.CodeBlock) any {
	e.frameStart()
	fe := e.new((tree.Ast)(fc.Nodes))
	_ = fe.run()
	e.frameEndWithRetAndBreak()
	return nil
}
