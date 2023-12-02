package eval

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
)

func (e *Eval) EvalStringLiteral(n *node.StringLiteral) string {
	return n.Value
}

func (e *Eval) EvalFloatLiteral(n *node.FloatLiteral) float64 {
	return n.Value
}

func (e *Eval) EvalIntLiteral(n *node.IntLiteral) int {
	return n.Value
}

func (e *Eval) EvalBoolLiteral(n *node.BoolLiteral) bool {
	return n.Value
}

func (e *Eval) EvalArrayLiteral(n *node.ArrayLiteral) []any {
	var arr []any
	for _, v := range n.Value {
		arr = append(arr, e.EvalExpr(v))
	}
	return arr
}

func (e *Eval) EvalMapLiteral(n *node.MapLiteral) map[any]any {
	var m = make(map[any]any)
	for _, v := range n.Value {
		m[e.EvalExpr(v.Key)] = e.EvalExpr(v.Value)
	}
	return m
}

func (e *Eval) EvalIdentifier(n *node.Identifier) any {
	v, ok := e.objTable.Get(n.Value)
	if ok {
		return v.Val
	}
	panic("No Var Found")
}

func (e *Eval) getZeroValue(t *node.Type) any {
	if t == nil {
		return nil
	}
	baseEval := e
	if t.Package != "" {
		newEval, ok := e.objTable.Get(t.Package)
		if ok {
			baseEval = newEval.Val.(*Eval)
		} else {
			panic("No Package Found")
		}
	}
	switch t.Name {
	case "int":
		return 0
	case "number", "float":
		return 0.0
	case "string":
		return ""
	case "bool":
		return false
	default:
		def, ok := getFromObjTable[*node.StructBlock](baseEval.objTable, t.Name)
		if !ok {
			panic("No Struct Definition Found")
		}
		var m = make(map[string]any)
		for variable, varDeclare := range def.Body {
			if varDeclare.Value != nil {
				m[variable] = baseEval.EvalExpr(varDeclare.Value)
				continue
			}
			m[variable] = baseEval.getZeroValue(varDeclare.Type)
		}
		return &obj.StructField{
			TypeAs: t,
			Fields: m,
		}
	}
}

func (e *Eval) EvalStructLiteral(n *node.StructLiteral) *obj.StructField {
	var sf *obj.StructField
	if n.Type != nil {
		sf = e.getZeroValue(n.Type).(*obj.StructField)
	} else {
		sf = &obj.StructField{
			Fields: make(map[string]any),
		}
	}

	for key, v := range n.Body {
		sf.Fields[key] = e.EvalExpr(v)
	}
	return sf
}
