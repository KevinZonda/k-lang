package idle

import (
	"github.com/gotk3/gotk3/gtk"
)

type MainW struct {
	*gtk.Window
	CodeEditor *CodeEditor
	MenuBar    *gtk.MenuBar
}

func NewMainW() *MainW {
	w := MainW{}
	w.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.MenuBar = NewMenuBar()
	w.SetTitle("IDLE Shell")
	w.SetDefaultSize(800, 600)

	w.CodeEditor = NewCodeEditor("cpp")

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)

	vbox.Add(w.MenuBar)
	vbox.PackStart(w.CodeEditor, true, true, 0)
	w.Add(vbox)

	return &w
}

func NewMenuBar() *gtk.MenuBar {
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
		exit, _ := gtk.MenuItemNewWithLabel("Exit")
		sep, _ := gtk.SeparatorMenuItemNew()
		_file.Append(newFile)
		_file.Append(openFile)
		_file.Append(sep)
		_file.Append(exit)
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
