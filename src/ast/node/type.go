package node

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/token"

type Type struct {
	Token    token.Token
	Package  string
	Name     string
	Array    bool
	Map      bool
	Func     bool
	Nullable bool
}

func (t *Type) CodeName() string {
	tail := ""
	if t.Nullable {
		tail = "?"
	}
	if t.Func {
		return "fn" + tail
	}
	if t.Array {
		return "any[]" + tail
	}
	if t.Map {
		return "map" + tail
	}
	tN := ""
	if t.Package != "" {
		tN += t.Package + "."
	}
	tN += t.Name
	return tN + tail

}
