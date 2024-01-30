package idle

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/fmtr"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle/gtks"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/gotk3/gotk3/gtk"
	"strconv"
	"strings"
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

	e    *eval.Eval
	Path string
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
	w.SetTitle("IDLE - " + w.Path)
}

func (w *EditorW) syncCursorPos() {
	buf := w.CodeE.Buf
	iter := buf.GetIterAtMark(buf.GetInsert())
	line := iter.GetLine()
	col := iter.GetLineOffset()
	w.StatusBar.Push(0, fmt.Sprintf("Line %d, Col %d", line+1, col+1))
}

func NewEditorW() *EditorW {
	w := EditorW{}
	w.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.SetTitle("IDLE")

	w.CodeE = gtks.NewCodeEditor("cpp")
	w.CodeView = gtks.WrapToScrolledWindow(w.CodeE)

	w.ReplE = gtks.NewCodeEditor("markdown")
	w.ReplE.SetEditable(false)
	w.ReplE.SetShowLineNumbers(false)
	w.ReplView = gtks.WrapToScrolledWindow(w.ReplE)

	w.Toolbar = w.NewToolBar()
	w.MenuBar = NewMenuBar()

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
	w.ReplEnter.Connect("activate", w.InvokeUserRepl)
	return &w
}

func (w *EditorW) RunCode() {
	w.runCode(true)
}

func (w *EditorW) RerunCode() {
	w.runCode(false)
}

func (w *EditorW) runCode(loadCtx bool) {
	w.ReplE.SmartNewLine()
	w.ReplE.AppendEnd("===================NEW RUN===================\n")
	ast, errs := parserHelper.Ast(w.CodeE.Text())
	if len(errs) > 0 {
		w.ReplE.AppendEnd(parseErrors(errs))
		return
	}
	ev := eval.New(ast, "")
	if loadCtx {
		ev.LoadContext(w.e)
	}
	w.e = ev
	isPanic, stdio, panicMsg, _, _ := runCode(ev)
	if !isPanic {
		w.ReplE.AppendEnd(stdio)
		w.ReplE.ScrollToEnd()
	} else {
		w.ReplE.AppendEnd(stdio)
		msg := stdio + "Code Panicked:\n" + panicMsg
		dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, msg)
		dialog.SetTitle("Result with Panic")
		dialog.Run()
		dialog.Destroy()
	}
}

func (w *EditorW) FormatCode() {
	code := w.CodeE.Text()
	w.CodeE.SetText(fmtr.Fmt(code))
}

func parseErrors(errs []parserHelper.CodeError) string {
	sb := strings.Builder{}
	sb.WriteString("Parse failed:\n")
	for idx, err := range errs {
		sb.WriteString("[" + strconv.Itoa(idx) + "] ")
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (w *EditorW) InvokeUserRepl() {
	cmd, _ := w.ReplEnter.GetText()
	w.ReplEnter.SetText("")

	w.ReplE.SmartNewLine()
	w.ReplE.AppendEnd(">>> " + cmd + "\n")
	ast, errs := parserHelper.Ast(cmd)
	if len(errs) > 0 {
		w.ReplE.AppendEnd(parseErrors(errs))
		return
	}
	ev := eval.New(ast, "")
	ev.LoadContext(w.e)
	w.e = ev
	isPanic, stdio, panicMsg, val, hasVal := runCode(ev)
	if !isPanic {
		w.ReplE.AppendEnd(stdio)
		if hasVal {
			w.ReplE.AppendEnd(fmt.Sprintf("<<< %v\n", val))
		}
		w.ReplE.ScrollToEnd()
	} else {
		msg := stdio + "\nBut with following panic:\n" + panicMsg
		dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, msg)
		dialog.SetTitle("Result with Panic")
		dialog.Run()
		dialog.Destroy()
	}
}

func runCode(e *eval.Eval) (isPanic bool, stdio string, panicMsg string, retV any, hasRet bool) {
	buf := NewFakeWCloser()
	defer func() {
		stdio = buf.ReadAllString()
		if r := recover(); r != nil {
			isPanic = true
			switch rT := r.(type) {
			case string:
				panicMsg = rT
			case error:
				panicMsg = rT.Error()
			default:
				panicMsg = "Unknown panic: " + fmt.Sprint(r)
			}

		}
	}()
	e.SetStdOut(buf)
	e.SetStdErr(buf)
	val, hasVal := e.DoRetLastExpr()

	return false, "", "", val, hasVal
}

type ToolBar struct {
	*gtk.Toolbar
	RestartBtn *gtk.ToolButton
	RunBtn     *gtk.ToolButton
	FmtBtn     *gtk.ToolButton
}

func (w *EditorW) NewToolBar() *ToolBar {
	bar := ToolBar{}
	bar.Toolbar, _ = gtk.ToolbarNew()
	bar.SetStyle(gtk.TOOLBAR_BOTH)

	bar.RunBtn, _ = gtk.ToolButtonNew(nil, "Run")
	bar.RunBtn.SetIconName("media-playback-start")

	bar.RestartBtn, _ = gtk.ToolButtonNew(nil, "Restart")
	bar.RestartBtn.SetIconName("view-refresh")

	bar.FmtBtn, _ = gtk.ToolButtonNew(nil, "Format")
	bar.FmtBtn.SetIconName("format-indent-more") //"document-page-setup")
	sep, _ := gtk.SeparatorToolItemNew()

	bar.Insert(bar.FmtBtn, 0)
	bar.Insert(sep, 1)
	bar.Insert(bar.RestartBtn, 2)
	bar.Insert(bar.RunBtn, 3)
	return &bar
}
