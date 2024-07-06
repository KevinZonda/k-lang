package symbol

type File struct {
	PublicSymbol  map[string]*Element
	PrivateSymbol map[string]*Element
}

// var AllMetaData map[string]*File = map[string]*File{}

func NewFile() *File {
	return &File{
		PublicSymbol:  map[string]*Element{},
		PrivateSymbol: map[string]*Element{},
	}
}

type Element struct {
	*File
	Kind ElementKind
	Type Type
	Keys []string
}

type Type []string

func (e *Element) AddChild(name string, kind ElementKind, t Type) {
	e.AddChildElem(name, &Element{Kind: kind, Type: t})
}

func (e *Element) AddChildElem(name string, elem *Element) {
	if e.File == nil {
		e.File = NewFile()
	}
	tgt := e.File.PubOrPriv(name)
	tgt[name] = elem
}

type ElementKind string

const (
	ElementKindStruct   ElementKind = "struct"
	ElementKindFunction ElementKind = "function"
	ElementKindVariable ElementKind = "variable"
)
