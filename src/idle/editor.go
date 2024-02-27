package idle

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/fmtr"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle/gtks"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/gotk3/gotk3/gtk"
	"io"
	"runtime"
	"sync"
)

type EditorW struct {
	*gtk.Window

	CodeE    *gtks.CodeEditor
	CodeView *gtk.ScrolledWindow

	MenuBar *gtk.MenuBar
	Toolbar *ToolBar
	VBox    *gtk.Box

	ReplE     *gtks.CodeEditor
	ReplView  *gtk.ScrolledWindow
	ReplEnter *gtk.Entry
	ReplVBox  *gtk.Box

	StatusBar *gtk.Statusbar

	e       *eval.Eval
	changed bool
	Path    string

	evalIn io.ReadWriter

	PanicWithDlg bool

	isRunning  bool
	cancelFunc func()
	gtkIO      *sync.Mutex

	custom Customizer

	acc *gtk.AccelGroup
}

func (w *EditorW) SetChanged(changed bool) {
	w.changed = changed
	w.syncTitle()
}

func (w *EditorW) syncTitle() {
	p := w.Path
	if p == "" {
		p = "Untitled"
	}
	if w.changed {
		w.SetTitle("IDLE - " + p + " *")
	} else {
		w.SetTitle("IDLE - " + p)
	}
}

type ReplETags struct {
	Red  *gtk.TextTag
	Blue *gtk.TextTag
}

func (w *EditorW) OpenFile() {
	fc, _ := gtk.FileChooserNativeDialogNew("Open file...", w, gtk.FILE_CHOOSER_ACTION_OPEN, "_Open", "_Cancel")
	resp := fc.NativeDialog.Run()
	if gtk.ResponseType(resp) == gtk.RESPONSE_ACCEPT {
		w.LoadFile(fc.GetFilename())
		w.ShowAll()
	}
}

func (w *EditorW) LoadFile(path string) {
	s, err := iox.ReadAllText(path)
	if err != nil {
		dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, err.Error())
		dialog.SetTitle("Open File Failed")
		dialog.Run()
		dialog.Destroy()
		return
	}
	w.Path = path
	w.CodeE.SetText(s)
	w.SetChanged(false)
}

func NewEditorW() *EditorW {
	w := EditorW{
		gtkIO: &sync.Mutex{},
	}
	customizeOk := w.onNewWindow()
	lifecycle.IncrementCount()
	w.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.syncTitle()

	w.acc, _ = gtk.AccelGroupNew()

	fontSizeInt := 16

	w.CodeE = gtks.NewCodeEditor("cpp", fontSizeInt)
	w.CodeView = gtks.WrapToScrolledWindow(w.CodeE)
	w.CodeE.OnChanged(func() {
		w.SetChanged(true)
	})

	w.ReplE = gtks.NewCodeEditor("", fontSizeInt)
	w.ReplE.SetEditable(false)
	w.ReplE.SetShowLineNumbers(false)
	w.ReplE.SetWrapMode(gtk.WRAP_CHAR)
	w.ReplE.Tags["red"] = w.ReplE.NewTextTag("red", "#ff0000")
	w.ReplE.Tags["blue"] = w.ReplE.NewTextTag("blue", "#0000ff")

	w.ReplView = gtks.WrapToScrolledWindow(w.ReplE)

	w.Toolbar = w.NewToolBar()
	w.MenuBar = w.NewMenuBar()

	w.StatusBar, _ = gtk.StatusbarNew()
	w.StatusBar.SetBorderWidth(0)
	w.StatusBar.SetMarginBottom(0)
	w.StatusBar.SetMarginTop(0)
	w.StatusBar.SetMarginStart(0)
	w.StatusBar.SetMarginEnd(0)

	w.ReplEnter, _ = gtk.EntryNew()
	w.ReplEnter.SetPlaceholderText("Enter REPL command here")
	w.ReplVBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	w.ReplVBox.PackStart(w.ReplView, true, true, 0)
	w.ReplVBox.Add(w.ReplEnter)

	contentPane, _ := gtk.PanedNew(gtk.ORIENTATION_VERTICAL)

	contentPane.Pack1(w.CodeView, true, false)
	contentPane.Pack2(w.ReplVBox, true, true)

	w.VBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if runtime.GOOS != "darwin" || !customizeOk {
		w.VBox.Add(w.MenuBar)
	} else {
		// darwin && customizeOk
		// MenuBar in macOS is handled by gtk_osxapplication
		// but not show will also cause the menu to be hidden
		w.MenuBar.ShowAll()
	}
	w.VBox.Add(w.Toolbar)
	w.VBox.PackStart(contentPane, true, true, 0)
	w.VBox.Add(w.StatusBar)

	w.Add(w.VBox)
	w.SetDefaultSize(800, 600)

	w.syncCursorPos()
	w.CodeE.SourceView.TextView.Connect("move-cursor", w.syncCursorPos)
	w.CodeE.Buf.Connect("end-user-action", w.syncCursorPos)
	w.ReplEnter.Connect("activate", w.InvokeUserRepl)

	w.shortcut(ctrl()+"r", w.RerunCode)
	w.shortcut(ctrl()+"s", w.Save)
	w.shortcut(ctrl()+"o", w.OpenFile)
	w.shortcut(ctrl()+"w", w.Close)

	w.AddAccelGroup(w.acc)

	w.Connect("destroy", func() {
		lifecycle.Decrease()
	})
	w.syncRunningStat()

	w.ReplE.AppendEnd(buildconst.Msg() + "\n")
	w.startPrompt()

	w.onWindowCreated()

	return &w
}

func (w *EditorW) Stop() {
	if !w.isRunning {
		return
	}
	if w.cancelFunc != nil {
		w.cancelFunc()
		w.setRunning(false)
		w.evalIn = nil
	}
}

func (w *EditorW) setRunning(v bool) {
	w.isRunning = v
	w.syncRunningStat()
}

func (w *EditorW) FormatCode() {
	code := w.CodeE.Text()
	code, err := fmtr.Fmt(code)
	if len(err) > 0 {
		w.ReplE.SmartNewLine()
		w.ReplE.AppendTag(w.ReplE.Tags["red"], err.String())
		w.ReplE.ScrollToEnd()
		return
	}
	w.CodeE.SetText(code)
}

func (w *EditorW) InvokeUserRepl() {
	cmd, _ := w.ReplEnter.GetText()
	w.ReplEnter.SetText("")

	if w.isRunning {
		if w.evalIn != nil {
			cmd = cmd + "\n"
			io.WriteString(w.evalIn, cmd)
			w.ReplE.AppendEnd(cmd)
		}
		return
	}

	w.ReplE.AppendEnd(cmd + "\n")

	rst := w.runCode(cmd, true, "")
	var val any
	if rst.HasReturn {
		val = rst.ReturnValue
	} else if rst.IsLastExpr {
		val = rst.LastExprVal
	} else {
		goto end
	}
	w.ReplE.SmartNewLine()
	w.ReplE.AppendTag(w.ReplE.Tags["blue"], "<<< ")
	w.ReplE.AppendEnd(fmt.Sprintf("%v\n", val))
end:
	w.startPrompt()
	w.ReplE.ScrollToEnd()
}

type ToolBar struct {
	*gtk.Toolbar
	StopBtn    *gtk.ToolButton
	SaveBtn    *gtk.ToolButton
	RestartBtn *gtk.ToolButton
	CleanBtn   *gtk.ToolButton
	RunBtn     *gtk.ToolButton
	FmtBtn     *gtk.ToolButton
}

func (w *EditorW) NewToolBar() *ToolBar {
	bar := ToolBar{}
	bar.Toolbar, _ = gtk.ToolbarNew()
	bar.SetStyle(gtk.TOOLBAR_BOTH)

	bar.RunBtn = gtks.ToolBtnWithIcon("Run", "media-playback-start", w.RunCode)
	bar.StopBtn = gtks.ToolBtnWithIcon("Stop", "media-playback-stop", w.Stop)
	bar.RestartBtn = gtks.ToolBtnWithIcon("Restart", "view-refresh", w.RerunCode)
	bar.CleanBtn = gtks.ToolBtnWithIcon("Clean", "edit-clear", func() {
		w.ReplE.SetText("")
		w.ReplE.AppendEnd(buildconst.Msg() + "\n")
		w.startPrompt()
	})
	bar.FmtBtn = gtks.ToolBtnWithIcon("Format", "format-indent-more", w.FormatCode)
	sep, _ := gtk.SeparatorToolItemNew()
	bar.SaveBtn = gtks.ToolBtnWithIcon("Save", "document-save", w.Save)

	items := []gtk.IToolItem{
		bar.SaveBtn, bar.FmtBtn, sep, bar.CleanBtn, bar.RestartBtn, bar.StopBtn, bar.RunBtn}
	for i, item := range items {
		bar.Toolbar.Insert(item, i)
	}
	return &bar
}

func (w *EditorW) SaveAs() {
	fc, _ := gtk.FileChooserNativeDialogNew("Save as file...", nil, gtk.FILE_CHOOSER_ACTION_SAVE, "_Save", "_Cancel")

	fc.AddFilter(gtks.NewFileFilter("*.k", "Klang source file"))
	fc.AddFilter(gtks.NewFileFilter("*.*", "All files"))

	resp := fc.NativeDialog.Run()
	if gtk.ResponseType(resp) == gtk.RESPONSE_ACCEPT {
		w.Path = fc.GetFilename()
		iox.WriteAllText(w.Path, w.CodeE.Text())
		w.SetChanged(false)
	}
}

func (w *EditorW) Save() {
	if w.Path == "" {
		w.SaveAs()
		return
	}
	iox.WriteAllText(w.Path, w.CodeE.Text())
	w.SetChanged(false)
}
