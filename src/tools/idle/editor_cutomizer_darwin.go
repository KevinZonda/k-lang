package idle

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/gotk3osx"
	"os"
)

type Customizer struct {
	osx *gotk3osx.GtkosxApplication
}

func (w *EditorW) onNewWindow() (ok bool) {
	if os.Getenv("K_LANG_GO_TEST") == "1" {
		return false
	}
	osx, err := gotk3osx.GetGtkosxApplication()
	if err != nil {
		osx = nil
	}
	w.custom = Customizer{osx: osx}
	return w.custom.osx != nil
}

func (w *EditorW) onWindowCreated() {
	if os.Getenv("K_LANG_GO_TEST") == "1" {
		return
	}
	if w.custom.osx == nil {
		return
	}
	w.custom.osx.SetMenuBar(&w.MenuBar.MenuShell)
	w.custom.osx.Ready()
}
