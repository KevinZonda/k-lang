package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) _evalIterArray(n *node.IterStyleFor, iters []any) any {
	for _, iter := range iters {
		_ = e.EvalLoopCodeBlockWithHook(n.Body, func() {
			e.objTable.SetAtTop(n.Variable.Value, iter)
		})
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
	return nil
}

func (e *Eval) _evalIterMap(n *node.IterStyleFor, iters map[any]any) any {
	for key, val := range iters {
		_ = e.EvalLoopCodeBlockWithHook(n.Body, func() {
			e.objTable.SetAtTop(n.Variable.Value, &obj.StructField{
				Fields: map[string]any{
					"key": key,
					"val": val,
				},
			})
		})
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
	return nil
}

func (e *Eval) EvalIterStyleForStmt(styleFor *node.IterStyleFor) any {
	e.loopLvl++
	iters := e.EvalExpr(styleFor.Iterator)
	switch iters.(type) {
	case []any:
		return e._evalIterArray(styleFor, iters.([]any))
	case map[any]any:
		return e._evalIterMap(styleFor, iters.(map[any]any))

	}
	panic("Not Supported Iteration Type")
}

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

func (e *Eval) EvalCStyleFrStmt(n *node.CStyleFor) any {
	e.loopLvl++
	if n.InitialExpr != nil {
		switch n.InitialExpr.(type) {
		case node.Stmt:
			e.EvalStmt(n.InitialExpr.(node.Stmt))
		case node.Expr:
			e.EvalExpr(n.InitialExpr.(node.Expr))
		}
	}
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
			e.EvalExpr(n.AfterIterExpr)
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
		e.EvalExpr(n.AfterIterExpr)
	}
}

func (e *Eval) EvalLoopCodeBlock(fc *node.CodeBlock) any {
	e.frameStart()
	fe := e.new((tree.Ast)(fc.Nodes))
	_ = fe.run()
	e.frameEndWithAll()
	return nil
}

func (e *Eval) EvalLoopCodeBlockWithHook(fc *node.CodeBlock, onNewFrame func()) any {
	e.frameStart()
	if onNewFrame != nil {
		onNewFrame()
	}
	fe := e.new((tree.Ast)(fc.Nodes))
	_ = fe.run()
	e.frameEndWithAll()
	return nil
}
