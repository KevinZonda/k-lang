package obj

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"strings"
)

type StructField struct {
	TypeAs     *node.Type
	Fields     *orderedmap.OrderedMap[string, any]
	ParentEval any
}

func NewStruct(parent any) *StructField {
	return &StructField{
		Fields:     orderedmap.New[string, any](),
		ParentEval: parent,
	}
}

func (s *StructField) With(key string, value any) *StructField {
	s.Fields.Set(key, value)
	return s
}

func (s *StructField) String() string {
	if VizAny {
		return s.Visualize()

	}
	return s.RawString()
}
func (s *StructField) RawString() string {
	if s == nil {
		return "<nil struct>"
	}
	sb := strings.Builder{}
	if s.TypeAs != nil {
		sb.WriteString(s.TypeAs.Name)
		sb.WriteString(" ")
	} else {
		sb.WriteString("struct ")
	}
	sb.WriteString("{")
	count := 0
	totalLen := s.Fields.Len()
	for pair := s.Fields.Oldest(); pair != nil; pair = pair.Next() {
		sb.WriteString(fmt.Sprint(pair.Key))
		sb.WriteString(": ")
		switch pairT := pair.Value.(type) {
		case *StructField:
			sb.WriteString(pairT.String())
		default:
			sb.WriteString(fmt.Sprint(pair.Value))
		}
		if count < totalLen-1 {
			sb.WriteString(", ")
		}
		count++
	}
	sb.WriteString("}")
	return sb.String()
}

func (s *StructField) Visualize() string {
	return TreeAnyW("struct", s, false).String()
}
