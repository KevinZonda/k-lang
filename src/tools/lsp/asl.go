package lsp

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/tree"
)

// Abstract Structure List
type ASL struct {
	funcs   map[string]*node.FuncBlock
	structs map[string]*node.StructBlock
	vars    map[string]*AslObj
	all     map[string]any
}

type AslObj struct {
	Name *node.Variable
	Type *node.Type
}

func astToASL(ast tree.Ast) *ASL {
	asl := &ASL{
		funcs:   map[string]*node.FuncBlock{},
		structs: map[string]*node.StructBlock{},
		all:     map[string]any{},
		vars:    map[string]*AslObj{},
	}
	for _, n := range ast {
		switch n.(type) {
		case *node.AssignStmt:
			for _, assignee := range n.(*node.AssignStmt).Assignee {
				v := &AslObj{
					Name: assignee.Var,
					Type: assignee.Type,
				}

				// FIXME: Could be a problem if there is package name?
				asl.vars[v.Name.Value[0].Name.Value] = v
				asl.all[v.Name.Value[0].Name.Value] = v

			}

		case *node.FuncBlock:
			fb := n.(*node.FuncBlock)
			asl.funcs[fb.Name.Value] = fb
			asl.all[fb.Name.Value] = fb
		case *node.StructBlock:
			structBlock := n.(*node.StructBlock)
			asl.structs[structBlock.Name] = structBlock
			asl.all[structBlock.Name] = structBlock
		}
	}
	return asl
}
