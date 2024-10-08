package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval/stringProcess"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
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
	tokens := stringProcess.Parse(mode, n.Value, n.Char)
	sb := strings.Builder{}
	for _, token := range tokens {
		switch token.Kind {
		case stringProcess.KindText:
			sb.WriteString(token.Value)
		case stringProcess.KindVar:
			tokenV := strings.TrimSpace(token.Value)
			if val, hasVal := e.memory.Get(tokenV); hasVal {
				sb.WriteString(fmt.Sprint(val.Value()))
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
		arr = append(arr, e.EvalExpr(v).EnsureValue())
	}
	return arr
}

func (e *Eval) EvalMapLiteral(n *node.MapLiteral) map[any]any {
	e.currentToken = n.GetToken()
	var m = make(map[any]any)
	for _, v := range n.Value {
		m[e.EvalExpr(v.Key).EnsureValue()] = e.EvalExpr(v.Value).EnsureValue()
	}
	return m
}

func (e *Eval) EvalIdentifier(n *node.Identifier, keepRef bool) any {
	e.currentToken = n.GetToken()
	v, ok := e.memory.Get(n.Value)
	if ok {
		if keepRef {
			return v
		}
		val := v.Value()
		switch valT := val.(type) {
		case *obj.Object:
			return valT.Value()
		default:
			return val
		}
	}
	if e.FeatUnknownVarNil {
		return nil
	}
	panic("No Var Found: " + n.Value)
}

func (e *Eval) getZeroValue(t *node.Type) any {
	if t == nil {
		return nil
	}
	if t.Nullable {
		return nil
	}
	if t.Array {
		return make([]any, 0)
	}
	if t.Map {
		return make(map[any]any, 0)
	}
	if t.Array {
		return nil
	}
	switch t.Name {
	case node.TypeInt:
		return 0
	case node.TypeNum:
		return 0.0
	case node.TypeString:
		return ""
	case node.TypeBool:
		return false
	case node.TypeAny:
		return nil
	default:
		baseEval := e.GetPackage(t.Package)
		def := getStructDef(baseEval, t, true)
		sf := obj.NewStruct(baseEval)
		for pair := def.Body.Oldest(); pair != nil; pair = pair.Next() {
			varName := pair.Key
			varDeclare := pair.Value
			sf.With(varName, baseEval.EvalVarDeclare(varDeclare))
		}
		sf.TypeAs = t
		return sf
	}
}

func (e *Eval) GetPackage(pkg string) *Eval {
	if pkg == "" {
		return e
	}
	p, ok := e.memory.Get(pkg)
	var ev *Eval
	if ok {
		ev = p.Value().(*Eval)
	}
	if !ok || ev == nil {
		panic("No Package Found: " + pkg)
	}
	return ev
}

func (e *Eval) EvalVarDeclare(varDeclare *node.Declare) any {
	if varDeclare.Func != nil {
		return varDeclare.Func
	}
	if varDeclare.Value != nil {
		return e.EvalExpr(varDeclare.Value).EnsureValue()
	}
	return e.getZeroValue(varDeclare.Type)
}

func (e *Eval) EvalStructLiteral(n *node.StructLiteral) *obj.StructField {
	e.currentToken = n.GetToken()
	var sf *obj.StructField
	if n.Type != nil {
		sf = e.getZeroValue(n.Type).(*obj.StructField)
	} else {
		sf = obj.NewStruct(e)
	}

	for pair := n.Body.Oldest(); pair != nil; pair = pair.Next() {
		sf.Fields.Set(pair.Key, e.EvalExpr(pair.Value).EnsureValue())
	}
	return sf
}
