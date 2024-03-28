package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"time"
)

type StdTimeLib struct {
	FBLibrary
}

func NewStdTimeLib() *StdTimeLib {
	return &StdTimeLib{
		FBLibrary: FBLibrary{
			F: map[string]*node.FuncBlock{
				"now":      FxToFuncBlock(time.Now().String),
				"unix":     FxToFuncBlock(i64txi(time.Now().Unix)),
				"unixNano": FxToFuncBlock(i64txi(time.Now().UnixNano)),
				"sleep":    FxToFuncBlock(func(a int) { time.Sleep(time.Duration(a) * time.Millisecond) }),
			},
		},
	}
}

func i64txi(f func() int64) func() int {
	return func() int {
		return int(f())
	}
}
