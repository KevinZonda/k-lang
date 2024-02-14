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
