package obj

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"

type StructField struct {
	TypeAs *node.Type
	Fields map[string]any
}
