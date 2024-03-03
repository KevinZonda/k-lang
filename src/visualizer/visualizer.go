package visualizer

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"github.com/xlab/treeprint"
	"reflect"
	"strings"
)

func VisualizeAny(name string, o any) *VNode {
	if o == nil {
		return &VNode{Val: nil}
	}

	switch vT := o.(type) {
	case int, float64, string, bool:
		return &VNode{Val: vT, Name: name}
	case []any:
		children := make([]*VNode, len(vT))
		for i, k := range vT {
			children[i] = VisualizeAny("", k)
		}
		return &VNode{Val: children, Name: name}
	case *obj.Object:
		return Visualize(name, vT)
	case *obj.StructField:
		f := vT.Fields
		children := map[string]*VNode{}
		for kvp := f.Oldest(); kvp != nil; kvp = kvp.Next() {
			vN := VisualizeAny(kvp.Key, kvp.Value)
			vN.Name = kvp.Key
			children[kvp.Key] = vN
		}
		return &VNode{Val: children, Name: name}
	default:
		return &VNode{Val: vT, Name: name}
	}
}

func TreeAny(t treeprint.Tree, name string, o any) treeprint.Tree {
	if o == nil {
		t.SetValue(name + ": nil")
		return t
	}

	switch vT := o.(type) {
	case int, float64, string, bool:
		t.SetValue(fmt.Sprintf("%s: %v", name, vT))
		return t
	case []any:
		t.SetValue(name + ": []")
		for idx, k := range vT {
			TreeAny(t.AddBranch(""), fmt.Sprint(idx), k)
		}
		return t
	case *obj.Object:
		return Tree(t, name, vT)
	case *obj.StructField:
		f := vT.Fields
		for kvp := f.Oldest(); kvp != nil; kvp = kvp.Next() {
			TreeAny(t.AddBranch(kvp.Key), kvp.Key, kvp.Value)
		}
		return t
	default:
		return nil
	}
}

func Tree(t treeprint.Tree, name string, o *obj.Object) treeprint.Tree {
	if t == nil {
		t = treeprint.New()
	}
	if o == nil {
		t.SetValue("nil")
		return t
	}

	t.SetValue(name)

	switch o.Kind {
	case obj.EvalObj:
		t.SetValue("Eval")
	case obj.Func:
		t.SetValue("Func")
	case obj.Lambda:
		t.SetValue("Lambda")
	case obj.Library:
		t.SetValue("Library")
	case obj.StructDef:
		t.SetValue("StructDef")
	case obj.Struct:
		TreeAny(t, name, o.Value())
	case obj.Value:
		t.AddNode(fmt.Sprintf("%v", o.Value()))
	}

	return t
}

func Visualize(name string, o *obj.Object) *VNode {
	if o == nil {
		return &VNode{Val: nil}
	}
	fmt.Println(reflect.TypeOf(o.Value()))

	switch o.Kind {
	case obj.EvalObj:
		return &VNode{Val: "Eval", Name: name}
	case obj.Func:
		return &VNode{Val: "Func", Name: o.ToFunc().Name.Value}
	case obj.Value:
		return VisualizeAny(name, o.Value())
	case obj.Lambda:
		return &VNode{Val: "Lambda"}
	case obj.Library:
		return &VNode{Val: "Library"}
	case obj.StructDef:
		return &VNode{Val: "StructDef", Name: o.ToStructDef().Name}
	case obj.Struct:
		return VisualizeAny(name, o.Value())
	}

	switch vT := o.Value().(type) {
	case int, float64, string, bool, []any:
		return VisualizeAny(name, vT)
	case *obj.Object:
		return Visualize(name, vT)
	case *obj.StructField:

	}
	return nil
}

type VNode struct {
	Name string
	Val  any
}

func (v *VNode) StringIdent(i int) string {
	sb := &strings.Builder{}
	fmt.Fprint(sb, ident(i))
	fmt.Fprint(sb, "name:", v.Name)
	fmt.Fprint(sb, ident(i))
	if v.Val == nil {
		fmt.Fprint(sb, "val : nil")
		return sb.String()
	}
	switch vT := v.Val.(type) {
	case int, float64, string, bool:
		fmt.Fprint(sb, "val :", vT)
	case []*VNode:
		fmt.Fprint(sb, "val :")
		for _, k := range vT {
			fmt.Fprint(sb, k.StringIdent(i+1))
		}
	case map[string]*VNode:
		fmt.Fprint(sb, "val :")
		for _, val := range vT {
			fmt.Fprint(sb, val.StringIdent(i+1))
		}
	case *VNode:
		fmt.Fprint(sb, "val :")
		fmt.Fprint(sb, vT.StringIdent(i+1))
	}

	return sb.String()
}

func ident(i int) string {
	s := "\n"
	for i > 0 {
		s += "    "
		i--
	}
	return s
}
