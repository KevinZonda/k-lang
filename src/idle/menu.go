package idle

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle/gtks"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func (w *EditorW) NewMenuBar() *gtk.MenuBar {
	mb, _ := gtk.MenuBarNew()

	fileMenu, _ := gtk.MenuItemNewWithLabel("File")
	{
		_file, _ := gtk.MenuNew()
		fileMenu.SetSubmenu(_file)
		newFile := gtks.NewMenuItem("New", func() {
			d := NewEditorW()
			d.ShowAll()
		})
		openFile := gtks.NewMenuItem("Open", func() {
			fc, _ := gtk.FileChooserNativeDialogNew("Open file...", nil, gtk.FILE_CHOOSER_ACTION_OPEN, "_Open", "_Cancel")
			resp := fc.NativeDialog.Run()
			if gtk.ResponseType(resp) == gtk.RESPONSE_ACCEPT {
				edit := NewEditorW()
				edit.LoadFile(fc.GetFilename())
				edit.ShowAll()
			}
		})
		saveFile, _ := gtk.MenuItemNewWithLabel("Save")
		saveFile.Connect("activate", w.Save)
		saveAsFile, _ := gtk.MenuItemNewWithLabel("Save As")
		saveAsFile.Connect("activate", w.SaveAs)

		exit, _ := gtk.MenuItemNewWithLabel("Exit")
		sep, _ := gtk.SeparatorMenuItemNew()
		_file.Append(newFile)
		_file.Append(openFile)
		_file.Append(saveFile)
		_file.Append(saveAsFile)
		_file.Append(sep)
		_file.Append(exit)
	}

	editMenu, _ := gtk.MenuItemNewWithLabel("Edit")
	{
		copyI := gtks.NewMenuItem("Copy", func() {
			clip, _ := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			w.CodeE.Buf.CopyClipboard(clip)
		})

		cutI := gtks.NewMenuItem("Cut", func() {
			clip, _ := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			w.CodeE.Buf.CutClipboard(clip, false)
		})

		pasteI := gtks.NewMenuItem("Paste", func() {
			clip, _ := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			w.CodeE.Buf.PasteClipboard(clip, nil, true)
		})

		delI := gtks.NewMenuItem("Delete", func() {
			w.CodeE.Buf.DeleteSelection(false, false)
		})

		clearI := gtks.NewMenuItem("Clear", func() {
			w.CodeE.SetText("")
		})

		_edit, _ := gtk.MenuNew()

		_edit.Append(copyI)
		_edit.Append(cutI)
		_edit.Append(pasteI)
		_edit.Append(delI)
		_edit.Append(clearI)

		editMenu.SetSubmenu(_edit)
	}

	helpItem, _ := gtk.MenuItemNewWithLabel("Help")
	{
		_help, _ := gtk.MenuNew()
		about := gtks.NewMenuItem("About", func() {
			d := HelpW()
			d.Run()
			d.Destroy()
		})
		_help.Append(about)

		helpItem.SetSubmenu(_help)
	}
	mb.Append(fileMenu)
	mb.Append(editMenu)
	mb.Append(helpItem)

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
	d.SetProgramName("IDLE")
	d.SetVersion("0.0.1")
	d.SetComments("IDLE (Integrated Development and Learning Environment) is a simple IDE for learning K programming language.")
	d.SetLicense("GUI based on GTK3 licensed under LGPL-2")
	return d
}
