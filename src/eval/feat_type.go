package eval

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"reflect"
)

func (e *Eval) TypeCheck(t *node.Type, v any) bool {
	if t == nil {
		return true
	}
	if t.Name == "any" {
		return true
	}
	if v == nil {
		return t.Nullable
	}
	switch v.(type) {
	case []any:
		return t.Package == "" && t.Array
	case map[any]any:
		return t.Package == "" && t.Map
	case int, float64:
		return t.IsPlainType(node.TypeInt) || t.IsPlainType(node.TypeNum)
	//case float64:
	//	return t.IsPlainType(node.TypeNum)
	case string:
		return t.IsPlainType(node.TypeString)
	case *node.LambdaExpr, *node.FuncBlock:
		return t.Func
	case *obj.StructField:
		return e.checkStructType(t, v.(*obj.StructField))
	case *obj.Object:
		vT := v.(*obj.Object)
		if vT.Is(obj.Func, obj.Lambda) {
			return t.Func
		}
		if vT.Is(obj.Value) {
			return e.TypeCheck(t, vT.Value())
		}
		if vT.Is(obj.Struct) {
			vStruct := vT.ToStruct()
			return e.TypeCheck(t, vStruct)
		}
		return true

	default:
		// TODO: more type check
		return true
	}
}

func getStructDef(baseEval *Eval, t *node.Type, ignorePack bool) *node.StructBlock {
	if !ignorePack {
		baseEval = baseEval.GetPackage(t.Package)
	}
	var def *node.StructBlock
	defObj, ok := baseEval.memory.Get(t.Name)
	if ok {
		def, ok = defObj.Value().(*node.StructBlock)
	}
	if !ok {
		panic("No Struct Definition Found: " + t.Name)
	}
	return def
}

func (e *Eval) getStructDef(t *node.Type) *node.StructBlock {
	return getStructDef(e, t, false)
}

func (e *Eval) checkStructType(t *node.Type, v *obj.StructField) bool {
	if v.TypeAs != nil && v.TypeAs.Name == t.Name {
		return true
	}

	// Only Check Type
	// TODO: NEED TYPE TX
	def := e.getStructDef(t)
	for pair := def.Body.Oldest(); pair != nil; pair = pair.Next() {
		varName := pair.Key
		varDeclare := pair.Value
		if varDeclare.Func != nil {
			continue
		}
		vField, fieldOk := v.Fields.Get(varName)
		if !fieldOk {
			return false
		}

		if !e.TypeCheck(varDeclare.Type, vField) {
			return false
		}
	}
	return true
}

func (e *Eval) NormaliseWithType(t *node.Type, v any) any {
	if t == nil {
		return v
	}
	if t.IsPlainType(node.TypeInt) {
		switch vT := v.(type) {
		case float64:
			return int(vT)
		case int:
			return vT
		default:
			return v
			//panic("Not Possible Type Transformation")
		}
	}
	if t.IsPlainType(node.TypeNum) {
		switch vT := v.(type) {
		case float64:
			return vT
		case int:
			return float64(vT)
		default:
			return v
			//panic("Not Possible Type Transformation")
		}
	}
	if vT, ok := v.(*obj.StructField); ok {
		if vT.TypeAs != nil && vT.TypeAs.Name == t.Name {
			return v
		}
		def := e.getStructDef(t)
		m := orderedmap.New[string, any]()
		for pair := def.Body.Oldest(); pair != nil; pair = pair.Next() {
			varName := pair.Key
			varDeclare := pair.Value
			if varDeclare.Func != nil {
				m.Set(varName, varDeclare.Func)
				continue
			}
			vField, fieldOk := vT.Fields.Get(varName)
			if !fieldOk {
				vField = e.getZeroValue(varDeclare.Type)
			}
			m.Set(varName, e.NormaliseWithType(varDeclare.Type, vField))
		}
		return &obj.StructField{
			TypeAs:     t,
			Fields:     m,
			ParentEval: vT.ParentEval,
		}
	}

	return v
}

func (e *Eval) TypeCast(t *node.Type, v any) any {
	if t == nil {
		return v
	}
	e.TypeCheckOrPanic(t, v)
	return e.NormaliseWithType(t, v)
}

func (e *Eval) TypeCheckOrPanic(t *node.Type, v any) {
	checked := e.TypeCheck(t, v)
	if !checked {
		panic(fmt.Sprintf("TypeCheck Failed, expected %s, got %s (%s)", t.CodeName(), reflect.TypeOf(v), fmt.Sprint(v)))
	}
}

func (e *Eval) AutoType(v any) *node.Type {
	switch v.(type) {
	case []any:
		return &node.Type{
			Array: true,
		}
	case map[any]any:
		return &node.Type{
			Map: true,
		}
	case *node.LambdaExpr, *node.FuncBlock:
		return &node.Type{
			Func: true,
		}
	case int:
		return &node.Type{
			Name: node.TypeInt,
		}
	case float64:
		return &node.Type{
			Name: node.TypeNum,
		}
	case string:
		return &node.Type{
			Name: node.TypeString,
		}
	case *obj.StructField:
		vT := v.(*obj.StructField)
		if vT == nil {
			return node.NewType()
		}
		return vT.TypeAs
	case *obj.Object:
		vT := v.(*obj.Object)
		if vT.Is(obj.Lambda, obj.Func) {
			return &node.Type{
				Func: true,
			}
		}
		if vT.Is(obj.Struct) {
			return e.AutoType(vT.ToStruct())
		}
	}
	// TODO: more type check
	// fmt.Println(reflect.TypeOf(v))
	return nil
}
