package idle

import (
	gtks2 "git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/idle/gtks"
	"github.com/gotk3/gotk3/gtk"
)

func (w *EditorW) NewMenuBar() *gtk.MenuBar {
	mb, _ := gtk.MenuBarNew()

	file := gtks2.NewMenu().
		AppendNewMenuItem("New", func() {
			d := NewEditorW()
			d.ShowAll()
		}).
		AppendNewMenuItem("Open", w.OpenFile).
		AppendNewMenuItem("Save", w.Save).
		AppendNewMenuItem("Save as", w.SaveAs).
		AppendSeparator().
		AppendNewMenuItem("Quit", w.Close)
	fileMenu, _ := gtk.MenuItemNewWithLabel("File")
	fileMenu.SetSubmenu(file)

	edit := gtks2.NewMenu().
		AppendNewMenuItem("Copy", func() {
			w.CodeE.Buf.CopyClipboard(gtks2.Clipboard())
		}).
		AppendNewMenuItem("Cut", func() {
			w.CodeE.Buf.CutClipboard(gtks2.Clipboard(), false)
		}).
		AppendNewMenuItem("Paste", func() {
			w.CodeE.Buf.PasteClipboard(gtks2.Clipboard(), nil, true)
		}).
		AppendNewMenuItem("Delete", func() {
			w.CodeE.Buf.DeleteSelection(false, false)
		}).AppendNewMenuItem("Clear", func() {
		w.CodeE.SetText("")
	})
	editMenu, _ := gtk.MenuItemNewWithLabel("Edit")
	editMenu.SetSubmenu(edit)

	help := gtks2.NewMenu().
		AppendNewMenuItem("About", func() {
			d := HelpW()
			d.Run()
			d.Destroy()
		})
	helpMenu, _ := gtk.MenuItemNewWithLabel("Help")
	helpMenu.SetSubmenu(help)

	mb.Append(fileMenu)
	mb.Append(editMenu)
	mb.Append(helpMenu)

	return mb
}

func ignoreErr[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func HelpW() *gtk.AboutDialog {
	d, _ := gtk.AboutDialogNew()
	d.SetProgramName("K Language IDLE")
	d.SetVersion("0.0.1")
	d.SetComments("IDLE (Integrated Development and Learning Environment) is a simple IDE for learning K programming language.")
	d.SetLicense("GUI based on GTK3 licensed under LGPL-2")
	return d
}
