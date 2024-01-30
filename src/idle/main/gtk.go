package main

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	//if runtime.GOOS == "windows" {
	//	settings, _ := gtk.SettingsGetDefault()
	//	settings.SetProperty("gtk-theme-name", "win32")
	//}

	w := idle.NewEditorW()
	w.Connect("destroy", func() {
		gtk.MainQuit()
	})
	w.ShowAll()

	gtk.Main()
}
