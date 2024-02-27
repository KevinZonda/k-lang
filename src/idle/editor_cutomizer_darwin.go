package idle

import "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/gotk3osx"

type Customizer struct {
	osx *gotk3osx.GtkosxApplication
}

func (w *EditorW) onNewWindow() {
	osx, err := gotk3osx.GetGtkosxApplication()
	if err != nil {
		osx = nil
	}
	w.custom = Customizer{osx: osx}
}

func (w *EditorW) onWindowCreated() {
	if w.custom.osx == nil {
		return
	}
	w.custom.osx.SetMenuBar(&w.MenuBar.MenuShell)
	w.custom.osx.Ready()
}
