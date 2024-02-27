package idle

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle/gtks"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type MainW struct {
	*gtk.Window
	CodeEditor *gtks.CodeEditor
	MenuBar    *gtk.MenuBar
}

func (w *EditorW) shortcut(acc string, f func()) {
	key, mod := gtk.AcceleratorParse(acc)
	w.acc.Connect(key, mod, gtk.ACCEL_VISIBLE, f)
}

func (w *EditorW) NewMenuBar() *gtk.MenuBar {
	mb, _ := gtk.MenuBarNew()

	fileMenu, _ := gtk.MenuItemNewWithLabel("File")
	{
		_file, _ := gtk.MenuNew()
		fileMenu.SetSubmenu(_file)
		newFile, _ := gtk.MenuItemNewWithLabel("New")

		newFile.Connect("activate", func() {
			d := NewEditorW()
			d.ShowAll()
		})
		openFile, _ := gtk.MenuItemNewWithLabel("Open")
		openFile.Connect("activate", func() {
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
		copyI, _ := gtk.MenuItemNewWithLabel("Copy")
		copyI.Connect("activate", func() {
			clip, _ := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			w.CodeE.Buf.CopyClipboard(clip)
		})

		cutI, _ := gtk.MenuItemNewWithLabel("Cut")
		cutI.Connect("activate", func() {
			clip, _ := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			w.CodeE.Buf.CutClipboard(clip, false)
		})

		pasteI, _ := gtk.MenuItemNewWithLabel("Paste")
		pasteI.Connect("activate", func() {
			clip, _ := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
			w.CodeE.Buf.PasteClipboard(clip, nil, true)
		})

		delI, _ := gtk.MenuItemNewWithLabel("Delete")
		delI.Connect("activate", func() {
			w.CodeE.Buf.DeleteSelection(false, false)
		})

		clearI, _ := gtk.MenuItemNewWithLabel("Clear")
		clearI.Connect("activate", func() {
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
		about, _ := gtk.MenuItemNewWithLabel("About")
		about.Connect("activate", func() {
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
