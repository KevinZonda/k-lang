package funcs

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/fmtr"
	"github.com/KevinZonda/GoX/pkg/iox"
)

func Format(path string) {
	txt, err := iox.ReadAllText(path)
	if err != nil {
		panic(err)
	}
	txt = fmtr.Fmt(txt)
	err = iox.WriteAllText(path, txt)
	if err != nil {
		panic(err)
	}
}
