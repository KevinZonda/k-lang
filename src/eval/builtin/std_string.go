package builtin

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/ast/node"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/obj"
	"strconv"
	"strings"
)

type StdStringLib struct {
	FBLibrary
}

func NewStdStringLib() *StdStringLib {
	return &StdStringLib{
		FBLibrary: FBLibrary{
			F: map[string]*node.FuncBlock{
				"len":       FxToFuncBlock(func(a string) int { return len([]rune(a)) }),
				"fromAscii": FxToFuncBlock(func(a int) string { return string(rune(a)) }),
				"trim":      FxToFuncBlock(strings.TrimSpace),
				"trimLeft":  FxToFuncBlock(func(a string) string { return strings.TrimLeft(a, " \t\n\r") }),
				"trimRight": FxToFuncBlock(func(a string) string { return strings.TrimRight(a, " \t\n\r") }),
				"split": FxToFuncBlock(func(a string, b string) []any {
					sps := strings.Split(a, b)
					var res []any
					for _, sp := range sps {
						res = append(res, sp)
					}
					return res
				}),
				"int": FxToFuncBlock(func(a string) int {
					i, e := strconv.ParseInt(a, 10, 64)
					if e != nil {
						panic(e)
					}
					return int(i)
				}),
				"num": FxToFuncBlock(func(a string) float64 {
					f, e := strconv.ParseFloat(a, 64)
					if e != nil {
						panic(e)
					}
					return f
				}),
			},
		},
	}
}

var _ obj.ILibrary = (*StdStringLib)(nil)
