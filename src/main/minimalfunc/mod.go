package minimalfunc

import (
	module2 "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/module"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
)

func Mod(arg string, args []string) {
	switch arg {
	case "new", "init":
		if len(args) != 1 {
			panic("expect mod new [console|lib]")
		}
		var m module2.Mod
		switch args[0] {
		case "console":
			m = module2.NewConsoleMod()
			iox.WriteAllText("main.k",
				`println('Hello World!')`)
		case "lib":
			m = module2.NewMod()
		default:
			panic("expect mod new [console|lib]")
		}

		err := iox.WriteAllText(module2.DEFAULT_K_MOD, m.String())
		panicx.PanicIfNotNil(err, err)

	case "restore", "download":
		mod := module2.LoadFromPath(module2.DEFAULT_K_MOD)
		mod.Restore()
	}
}
