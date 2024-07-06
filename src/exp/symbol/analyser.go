package symbol

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
	"strings"
)

func AnalyseAst(ast tree.Ast) *File {
	f := NewFile()
	for _, n := range ast {
		switch nT := n.(type) {
		case *node.CodeBlock:
			// No need to do
		case node.Expr:
			// No need to do
		case *node.AssignStmt:
			assignees := nT.Assignee
			for _, assignee := range assignees {
				if assignee.Var == nil {
					continue
				}
				names := assignee.Var.Value
				if len(names) != 1 {
					continue
				}
				name := names[0].Name.Value
				sym := f.PubOrPriv(name)
				sym[name] = &Element{
					Type: typeAsType(assignee.Type),
					Kind: ElementKindVariable,
				}
			}

		case *node.FuncBlock:
			fxName := nT.Name.Value
			sym := f.PubOrPriv(fxName)
			sym[fxName] = AnalyseFuncBlock(nT)
		case *node.StructBlock:
			name := nT.Name
			sym := f.PubOrPriv(name)
			sym[name] = AnalyseStructBlock(nT)
		default:
			// No need to do
		}
	}
	return f
}

func AnalyseFuncBlock(n *node.FuncBlock) *Element {
	var ts []string
	var sig []string
	for _, arg := range n.Args {
		if arg.Type != nil {
			t := arg.Type
			ts = append(ts, typeAsStr(t))
		} else {
			ts = append(ts, "any")
		}
		sig = append(sig, arg.Name.Value)
	}
	return &Element{
		Kind: ElementKindFunction,
		Type: ts,
		Keys: sig,
	}
}

func typeAsStr(t *node.Type) string {
	if t == nil {
		return ""
	}
	typeName := ""
	if t.Map {
		typeName = "map"
	} else if t.Array {
		typeName = "any[]"
	} else if t.Func {
		typeName = "fn"
	} else {
		if t.Package != "" {
			typeName = t.Package + "."
		}
		typeName += t.Name
	}
	if t.Nullable {
		typeName += "?"
	}
	return typeName
}

func typeAsType(t *node.Type) Type {
	if t == nil {
		return nil
	}
	typeName := ""
	if t.Map {
		typeName = "map"
	} else if t.Array {
		typeName = "any[]"
	} else if t.Func {
		typeName = "fn"
	} else {
		if t.Package != "" {
			typeName = t.Package + "."
		}
		typeName += t.Name
	}
	return []string{typeName}
}

func AnalyseStructBlock(n *node.StructBlock) *Element {
	e := &Element{
		Kind: ElementKindStruct,
	}
	for pair := n.Body.Oldest(); pair != nil; pair = pair.Next() {
		key := pair.Key
		val := pair.Value
		if val.Func != nil {
			e.AddChildElem(key, AnalyseFuncBlock(val.Func))
		} else {
			e.AddChild(key, ElementKindVariable, typeAsType(val.Type))
		}
	}
	return e
}

func (f *File) PubOrPriv(name string) map[string]*Element {
	if strings.HasPrefix(name, "_") {
		return f.PrivateSymbol
	}
	return f.PublicSymbol
}

func (f *File) Get(name string) *Element {
	return f.PubOrPriv(name)[name]
}
