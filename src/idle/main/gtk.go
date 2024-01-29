package main

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	editW := idle.NewEditorW()
	editW.Connect("destroy", func() {
		gtk.MainQuit()
	})
	editW.ShowAll()

	gtk.Main()
}
