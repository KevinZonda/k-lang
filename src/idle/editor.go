package idle

import (
	"bytes"
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/fmtr"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle/gtks"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"io"
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
	ReplETags *ReplETags
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

func (w *EditorW) syncCursorPos() {
	buf := w.CodeE.Buf
	iter := buf.GetIterAtMark(buf.GetInsert())
	line := iter.GetLine()
	col := iter.GetLineOffset()
	w.StatusBar.Push(0, fmt.Sprintf("Line %d, Col %d", line+1, col+1))
}

func (w *EditorW) syncRunningStat() {
	if !w.isRunning {
		w.cancelFunc = nil
	}
	glib.IdleAdd(func() {
		w.Toolbar.RunBtn.SetSensitive(!w.isRunning)
		w.Toolbar.StopBtn.SetSensitive(w.isRunning)
	})
}

func NewEditorW() *EditorW {
	w := EditorW{
		gtkIO: &sync.Mutex{},
	}
	w.onNewWindow()
	lifecycle.IncrementCount()
	w.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.syncTitle()

	fontSizeInt := 16

	w.CodeE = gtks.NewCodeEditor("cpp", fontSizeInt)
	w.CodeView = gtks.WrapToScrolledWindow(w.CodeE)
	{
		_buf, _ := w.CodeE.TextView.GetBuffer()
		_buf.Connect("changed", func() {
			w.SetChanged(true)
		})
	}

	w.ReplE = gtks.NewCodeEditor("", fontSizeInt)
	w.ReplE.SetEditable(false)
	w.ReplE.SetShowLineNumbers(false)
	w.ReplETags = &ReplETags{}
	w.ReplETags.Red = w.ReplE.NewTextTag("red", "#ff0000")
	w.ReplETags.Blue = w.ReplE.NewTextTag("blue", "#0000ff")

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
	w.VBox.Add(w.MenuBar)
	w.VBox.Add(w.Toolbar)
	w.VBox.PackStart(contentPane, true, true, 0)
	w.VBox.Add(w.StatusBar)

	w.Add(w.VBox)
	w.SetDefaultSize(800, 600)

	w.syncCursorPos()
	w.CodeE.SourceView.TextView.Connect("move-cursor", w.syncCursorPos)
	w.CodeE.Buf.Connect("end-user-action", w.syncCursorPos)

	w.Toolbar.RunBtn.Connect("clicked", w.RunCode)
	w.Toolbar.RestartBtn.Connect("clicked", w.RerunCode)
	w.Toolbar.FmtBtn.Connect("clicked", w.FormatCode)
	w.Toolbar.CleanBtn.Connect("clicked", func() {
		w.ReplE.SetText("")
		w.ReplE.AppendEnd(buildconst.Msg() + "\n")
		w.startPrompt()
	})
	w.ReplEnter.Connect("activate", w.InvokeUserRepl)

	{
		accel, _ := gtk.AccelGroupNew()
		key, mod := gtk.AcceleratorParse("<Control>s")
		accel.Connect(key, mod, gtk.ACCEL_VISIBLE, w.Save)
		key, mod = gtk.AcceleratorParse("<Control>o")
		accel.Connect(key, mod, gtk.ACCEL_VISIBLE, w.OpenFile)
		w.AddAccelGroup(accel)
	}

	w.Connect("destroy", func() {
		lifecycle.Decrease()
	})
	w.syncRunningStat()

	w.ReplE.AppendEnd(buildconst.Msg() + "\n")
	w.startPrompt()

	w.onWindowCreated()

	return &w
}

func (w *EditorW) RunCode() {
	go w.runCode(w.CodeE.Text(), true, "===================NEW RUN===================")
}

func (w *EditorW) RerunCode() {
	w.Stop()
	go w.runCode(w.CodeE.Text(), false, "===================NEW RUN===================")

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

func (w *EditorW) runCode(code string, loadCtx bool, beginMsg string) (rst eval.DetailedRunResult) {
	w.evalIn = &FakeStdIn{buf: &bytes.Buffer{}}
	if w.isRunning {
		return
	}

	defer func() {
		w.evalIn = nil
		w.setRunning(false)
	}()

	w.ReplE.SmartNewLine()
	if beginMsg != "" {
		w.ReplE.AppendEnd(beginMsg + "\n")
	}

	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		w.ReplE.AppendTag(w.ReplETags.Red, "Parse failed:\n"+errs.String())
		return
	}
	if !loadCtx || w.e == nil {
		w.e = eval.New()
	}
	ev := w.e
	stdout := w.ReplE.WriterPipe(w.gtkIO)
	ev.SetStdIn(w.evalIn)
	ev.SetStdOut(stdout)
	ev.SetStdErr(stdout)

	ch := make(chan eval.DetailedRunResult, 1)

	cancel := AsyncFunc(func() {
		ch <- ev.DoSafely(ast)
	})
	w.cancelFunc = func() {
		w.gtkIO.Lock()
		defer w.gtkIO.Unlock()
		cancel()
	}
	w.setRunning(true)
	rst = <-ch
	close(ch)
	glib.IdleAdd(func() {
		w.ReplE.ScrollToEnd()
		if rst.IsPanic {
			if w.PanicWithDlg {
				msg := "Code Panicked:\n" + rst.PanicMsg
				dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, msg)
				dialog.SetTitle("Result with Panic")
				dialog.Run()
				dialog.Destroy()
			} else {
				w.ReplE.AppendTagUnsafe(w.ReplETags.Red, "[PANIC RECEIVED] "+rst.PanicMsg+"\n")
				w.ReplE.ScrollToEnd()
			}
		}
		if beginMsg != "" { // User Invoke Will Handle
			w.startPrompt()
		}
	})
	return
}

func (w *EditorW) FormatCode() {
	code := w.CodeE.Text()
	code, err := fmtr.Fmt(code)
	if len(err) > 0 {
		w.ReplE.SmartNewLine()
		w.ReplE.AppendTag(w.ReplETags.Red, "Formatter Failed:\n"+err.String())
		w.ReplE.ScrollToEnd()
		return
	}
	w.CodeE.SetText(code)
}

func (w *EditorW) startPrompt() {
	w.ReplE.SmartNewLine()
	w.ReplE.AppendTag(w.ReplETags.Blue, ">>> ")
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
	w.ReplE.AppendTag(w.ReplETags.Blue, "<<< ")
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

	bar.RunBtn, _ = gtk.ToolButtonNew(nil, "Run")
	bar.RunBtn.SetIconName("media-playback-start")

	bar.StopBtn, _ = gtk.ToolButtonNew(nil, "Stop")
	bar.StopBtn.SetIconName("media-playback-stop")
	bar.StopBtn.Connect("clicked", w.Stop)

	bar.RestartBtn, _ = gtk.ToolButtonNew(nil, "Restart")
	bar.RestartBtn.SetIconName("view-refresh")

	bar.CleanBtn, _ = gtk.ToolButtonNew(nil, "Clean")
	bar.CleanBtn.SetIconName("edit-clear")

	bar.FmtBtn, _ = gtk.ToolButtonNew(nil, "Format")
	bar.FmtBtn.SetIconName("format-indent-more") //"document-page-setup")
	sep, _ := gtk.SeparatorToolItemNew()

	bar.SaveBtn, _ = gtk.ToolButtonNew(nil, "Save")
	bar.SaveBtn.SetIconName("document-save")
	bar.SaveBtn.Connect("clicked", w.Save)

	items := []gtk.IToolItem{
		bar.SaveBtn, bar.FmtBtn, sep, bar.CleanBtn, bar.RestartBtn, bar.StopBtn, bar.RunBtn}
	for i, item := range items {
		bar.Toolbar.Insert(item, i)
	}
	return &bar
}

func (w *EditorW) SaveAs() {
	fc, _ := gtk.FileChooserNativeDialogNew("Save as file...", nil, gtk.FILE_CHOOSER_ACTION_SAVE, "_Save", "_Cancel")
	{
		filter, _ := gtk.FileFilterNew()
		filter.AddPattern("*.k")
		filter.SetName("Klang source file")
		fc.AddFilter(filter)
	}
	{
		filter, _ := gtk.FileFilterNew()
		filter.AddPattern("*.*")
		filter.SetName("All files")
		fc.AddFilter(filter)
	}
	resp := fc.NativeDialog.Run()
	if gtk.ResponseType(resp) == gtk.RESPONSE_ACCEPT {
		w.Path = fc.GetFilename()
		iox.WriteAllText(w.Path, w.CodeE.Text())
		w.SetTitle("IDLE - " + w.Path)
	}

}

func (w *EditorW) Save() {
	if w.Path == "" {
		w.SaveAs()
		return
	}
	iox.WriteAllText(w.Path, w.CodeE.Text())
}
