package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/stringProcess"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"strings"
)

func (e *Eval) EvalStringLiteral(n *node.StringLiteral) string {
	e.currentToken = n.GetToken()
	mode := stringProcess.ModeNormal
	switch n.Mode {
	case '@':
		mode = stringProcess.ModeConst
	case '$':
		mode = stringProcess.ModeVar
	}
	tokens := stringProcess.Parse(mode, n.Value)
	sb := strings.Builder{}
	for _, token := range tokens {
		switch token.Kind {
		case stringProcess.KindText:
			sb.WriteString(token.Value)
		case stringProcess.KindVar:
			tokenV := strings.TrimSpace(token.Value)
			if val, hasVal := e.objTable.Get(tokenV); hasVal {
				sb.WriteString(fmt.Sprint(val.Val))
			} else {
				sb.WriteString("<!!! VAR NOT FOUND! = {" + tokenV + "}>")
			}
		}
	}
	return sb.String()
}

func (e *Eval) EvalFloatLiteral(n *node.FloatLiteral) float64 {
	e.currentToken = n.GetToken()
	return n.Value
}

func (e *Eval) EvalIntLiteral(n *node.IntLiteral) int {
	e.currentToken = n.GetToken()
	return n.Value
}

func (e *Eval) EvalBoolLiteral(n *node.BoolLiteral) bool {
	e.currentToken = n.GetToken()
	return n.Value
}

func (e *Eval) EvalArrayLiteral(n *node.ArrayLiteral) []any {
	e.currentToken = n.GetToken()
	var arr []any
	for _, v := range n.Value {
		arr = append(arr, e.EvalExpr(v))
	}
	return arr
}

func (e *Eval) EvalMapLiteral(n *node.MapLiteral) map[any]any {
	e.currentToken = n.GetToken()
	var m = make(map[any]any)
	for _, v := range n.Value {
		m[e.EvalExpr(v.Key)] = e.EvalExpr(v.Value)
	}
	return m
}

func (e *Eval) EvalIdentifier(n *node.Identifier) any {
	e.currentToken = n.GetToken()
	v, ok := e.objTable.Get(n.Value)
	if ok {
		return v.Val
	}
	panic("No Var Found: " + n.Value)
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
			panic("No Package Found: " + t.Package)
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
			panic("No Struct Definition Found: " + t.Name)
		}
		m := orderedmap.New[string, any]()
		// TODO: Is duplicate code??
		for pair := def.Body.Oldest(); pair != nil; pair = pair.Next() {
			variable := pair.Key
			varDeclare := pair.Value
			if varDeclare.Value != nil {
				m.Set(variable, baseEval.EvalExpr(varDeclare.Value))
				continue
			}
			m.Set(variable, baseEval.getZeroValue(varDeclare.Type))
		}
		return &obj.StructField{
			TypeAs: t,
			Fields: m,
		}
	}
}

func (e *Eval) EvalStructLiteral(n *node.StructLiteral) *obj.StructField {
	e.currentToken = n.GetToken()
	var sf *obj.StructField
	if n.Type != nil {
		sf = e.getZeroValue(n.Type).(*obj.StructField)
	} else {
		sf = &obj.StructField{
			Fields: orderedmap.New[string, any](),
		}
	}

	for pair := n.Body.Oldest(); pair != nil; pair = pair.Next() {
		sf.Fields.Set(pair.Key, e.EvalExpr(pair.Value))
	}
	return sf
}
