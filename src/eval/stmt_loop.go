package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/reserved"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) evalIterArray(n *node.IterStyleFor, iters []any) {
	e.currentToken = n.GetToken()
	loop := e.loopFrame()
	defer e.frameEnd()

	for _, iter := range iters {
		loop.SetValue(n.Variable.Value, iter)
		isContinue, isBreak, isReturn := e.loopDelegate(n.Body)
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
	}
	return
}

func (e *Eval) loopDelegate(body *node.CodeBlock) (isContinue, isBreak, isReturn bool) {
	e.EvalCodeBlock(body)
	top := e.memory.Top()
	if top.Has(reserved.Continue) {
		top.Remove(reserved.Continue)
		isContinue = true
		return
	}
	if top.Has(reserved.Return) {
		isReturn = true
		return
	}
	if top.Has(reserved.Break) {
		top.Remove(reserved.Break)
		isBreak = true
		return
	}
	return
}

func (e *Eval) evalIterMap(n *node.IterStyleFor, iters map[any]any) {
	e.currentToken = n.GetToken()
	loop := e.loopFrame()
	defer e.frameEnd()

	for key, val := range iters {
		loop.SetValue(n.Variable.Value,
			obj.NewStruct(nil).
				With("key", key).
				With("val", val))
		isContinue, isBreak, isReturn := e.loopDelegate(n.Body)
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
	}
	return
}

func (e *Eval) EvalIterStyleForStmt(styleFor *node.IterStyleFor) {
	e.currentToken = styleFor.GetToken()

	iters := e.EvalExpr(styleFor.Iterator).EnsureValue()
	switch itersT := iters.(type) {
	case string:
		// str to []rune to []string
		rs := []rune(itersT)
		as := make([]any, len(rs))
		for i, r := range rs {
			as[i] = string(rune(r))
		}
		e.evalIterArray(styleFor, as)
		return
	case []any:
		e.evalIterArray(styleFor, itersT)
		return
	case map[any]any:
		e.evalIterMap(styleFor, itersT)
		return
	}
	panic("Not Supported Iteration Type")
}

func (e *Eval) EvalWhileForStmt(n *node.WhileStyleFor) {
	e.currentToken = n.GetToken()

	e.loopFrame()
	defer e.frameEnd()

	for {
		if n.ConditionExpr != nil {
			rst := e.EvalExpr(n.ConditionExpr).EnsureValue().(bool)
			if !rst {
				return
			}
		}
		isContinue, isBreak, isReturn := e.loopDelegate(n.Body)
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
	}
}

func (e *Eval) EvalCStyleFrStmt(n *node.CStyleFor) {
	e.currentToken = n.GetToken()

	e.loopFrame()
	defer e.frameEnd()

	if n.InitialExpr != nil {
		switch initExpr := n.InitialExpr.(type) {
		case node.Stmt:
			e.EvalStmt(initExpr)
		case node.Expr:
			e.EvalExpr(initExpr)
		}
	}
	for {
		if n.ConditionExpr != nil {
			rst := e.EvalExpr(n.ConditionExpr).EnsureValue().(bool)
			if !rst {
				return
			}
		}
		isContinue, isBreak, isReturn := e.loopDelegate(n.Body)
		if isContinue {
			e.EvalExpr(n.AfterIterExpr)
		}
		if isContinue {
			continue
		}
		if isReturn || isBreak {
			return
		}
		e.EvalExpr(n.AfterIterExpr)
	}
}
