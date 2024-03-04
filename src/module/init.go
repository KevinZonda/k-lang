package module

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"

func NewMod() Mod {
	return Mod{
		KLangVersion: buildconst.LanguageVersion,
	}
}

func NewConsoleMod() Mod {
	return Mod{
		KLangVersion: buildconst.LanguageVersion,
		Entry:        DEFAULT_ENTRY,
	}
}
