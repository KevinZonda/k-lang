package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
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
		if e.objTable.HasKeyAtTop(reserved.Continue) {
			e.objTable.RemoveKeyAtTop(reserved.Continue)
			continue
		}
		if e.objTable.HasKeyAtTop(reserved.Return) {
			e.loopLvl--
			return nil
		}
		if e.objTable.HasKeyAtTop(reserved.Break) {
			e.loopLvl--
			e.objTable.RemoveKeyAtTop(reserved.Break)
			return nil
		}
	}
}

func (e *Eval) EvalLoopCodeBlock(fc *node.CodeBlock) any {
	e.frameStart()
	fe := e.new((tree.Ast)(fc.Nodes))
	_ = fe.run()
	e.frameEndWithRetBreakContinue()
	return nil
}
