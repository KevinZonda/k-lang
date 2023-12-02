package obj

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type StructField struct {
	TypeAs *node.Type
	Fields *orderedmap.OrderedMap[string, any]
}
