package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
)

func (e *Eval) EvalCommaExpr(c *node.CommaExpr) []any {
	return e.evalExprs(c.Exprs...)
}
