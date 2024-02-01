package main

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle"
	"github.com/gotk3/gotk3/gtk"
	"os"
)

func main() {
	gtk.Init(nil)

	//if runtime.GOOS == "windows" {
	//	settings, _ := gtk.SettingsGetDefault()
	//	settings.SetProperty("gtk-theme-name", "win32")
	//}
	w := idle.NewEditorW()

	if len(os.Args) > 1 {
		w.LoadFile(os.Args[1])
	}

	w.ShowAll()

	gtk.Main()
}
