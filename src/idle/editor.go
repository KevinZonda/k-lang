package idle

import (
	"github.com/gotk3/gotk3/gtk"
	sourceview "github.com/linuxerwang/sourceview3"
)

type EditorW struct {
	*gtk.Window
	CodeE    *CodeEditor
	Toolbar  *gtk.Toolbar
	CodeView *gtk.ScrolledWindow
	VBox     *gtk.Box
}

func NewEditorW() *EditorW {
	w := EditorW{}
	w.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.SetTitle("IDLE")

	w.VBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)

	w.CodeView, _ = gtk.ScrolledWindowNew(nil, nil)
	w.CodeView.SetVExpand(true)
	w.CodeE = NewCodeEditor("cpp")
	w.CodeView.Add(w.CodeE)

	w.Toolbar = NewToolBar()

	w.VBox.Add(w.Toolbar)
	w.VBox.Add(w.CodeView)

	w.Add(w.VBox)
	w.SetDefaultSize(800, 600)
	return &w
}

type CodeEditor struct {
	*sourceview.SourceView
	LanguageManager *sourceview.SourceLanguageManager
}

func NewCodeEditor(lang string) *CodeEditor {
	sv, _ := sourceview.SourceViewNew()
	lm, _ := sourceview.SourceLanguageManagerGetDefault()
	ce := CodeEditor{
		SourceView:      sv,
		LanguageManager: lm,
	}
	l, _ := lm.GetLanguage(lang)
	sv.SetShowLineNumbers(true)
	buf, _ := sv.GetBuffer()
	buf.SetLanguage(l)
	return &ce
}

func NewToolBar() *gtk.Toolbar {
	bar, _ := gtk.ToolbarNew()
	bar.SetStyle(gtk.TOOLBAR_BOTH)

	runBtn, _ := gtk.ToolButtonNew(nil, "Run")
	runBtn.SetIconName("media-playback-start")
	runBtn.Connect("clicked", func() {
		// Show a MessageBox
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Hello world!")
		dialog.Run()
		dialog.Destroy()

	})

	sep, _ := gtk.SeparatorToolItemNew()
	bar.Insert(sep, 0)
	bar.Insert(runBtn, 1)
	return bar
}
