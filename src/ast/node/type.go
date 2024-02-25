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

func (t *Type) plainType() bool {
	if t.Package != "" || t.Array || t.Map || t.Func {
		return false
	}
	return true
}

func (t *Type) IsPlainType(typeName string) bool {
	if !t.plainType() {
		return false
	}
	return t.Name == typeName
}

const (
	TypeInt    string = "int"
	TypeNum    string = "num"
	TypeString string = "string"
	TypeBool   string = "bool"
	TypeVoid   string = "void"
	TypeAny    string = "any"
)

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
