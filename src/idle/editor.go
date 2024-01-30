package idle

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/fmtr"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/idle/gtks"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/gotk3/gotk3/gtk"
	"log"
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

	ReplE    *gtks.CodeEditor
	ReplView *gtk.ScrolledWindow

	Path      string
	StatusBar *gtk.Statusbar
}

func (e *EditorW) LoadFile(path string) {
	s, err := iox.ReadAllText(path)
	if err != nil {
		dialog := gtk.MessageDialogNew(e, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, err.Error())
		dialog.SetTitle("Open File Failed")
		dialog.Run()
		dialog.Destroy()
		return
	}
	e.Path = path
	e.CodeE.SetText(s)
	e.SetTitle("IDLE - " + e.Path)
}

func (e *EditorW) syncCursorPos() {
	buf := e.CodeE.Buf
	iter := buf.GetIterAtMark(buf.GetInsert())
	line := iter.GetLine()
	col := iter.GetLineOffset()
	e.StatusBar.Push(0, fmt.Sprintf("Line %d, Col %d", line+1, col+1))
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

	contentPane, _ := gtk.PanedNew(gtk.ORIENTATION_VERTICAL)

	contentPane.Pack1(w.CodeView, true, false)
	contentPane.Pack2(w.ReplView, true, false)

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
	w.Toolbar.FmtBtn.Connect("clicked", w.FormatCode)

	return &w
}

func (e *EditorW) RunCode() {
	log.Println("RunCode")
	e.ReplE.AppendEnd("\n===================NEW RUN===================\n")
	ast, errs := parserHelper.Ast(e.CodeE.Text())
	if len(errs) > 0 {
		sb := strings.Builder{}
		sb.WriteString("Parse failed:\n")
		for idx, err := range errs {
			sb.WriteString("[" + strconv.Itoa(idx) + "] ")
			sb.WriteString(err.Error())
			sb.WriteString("\n")
		}
		e.ReplE.AppendEnd(sb.String())
		//dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, strings.TrimSpace(sb.String()))
		//dialog.SetTitle("Parse Failed")
		//dialog.Run()
		//dialog.Destroy()
		return
	}
	ev := eval.New(ast, "")
	isPanic, stdio, panicMsg := runCode(ev)
	if !isPanic {
		//dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, stdio)
		//dialog.SetTitle("Result")
		//dialog.Run()
		//dialog.Destroy()
		e.ReplE.AppendEnd(stdio)
	} else {
		msg := stdio + "\nBut with following panic:\n" + panicMsg
		dialog := gtk.MessageDialogNew(e, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, msg)
		dialog.SetTitle("Result with Panic")
		dialog.Run()
		dialog.Destroy()
	}
}

func (e *EditorW) FormatCode() {
	code := e.CodeE.Text()
	e.CodeE.SetText(fmtr.Fmt(code))
}

func runCode(e *eval.Eval) (isPanic bool, stdio string, panicMsg string) {
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
	e.Do()

	return false, "", ""
}

type ToolBar struct {
	*gtk.Toolbar
	RunBtn *gtk.ToolButton
	FmtBtn *gtk.ToolButton
}

func (e *EditorW) NewToolBar() *ToolBar {
	bar := ToolBar{}
	bar.Toolbar, _ = gtk.ToolbarNew()
	bar.SetStyle(gtk.TOOLBAR_BOTH)

	bar.RunBtn, _ = gtk.ToolButtonNew(nil, "Run")
	bar.RunBtn.SetIconName("media-playback-start")

	bar.FmtBtn, _ = gtk.ToolButtonNew(nil, "Format")
	bar.FmtBtn.SetIconName("format-indent-more") //"document-page-setup")
	sep, _ := gtk.SeparatorToolItemNew()

	bar.Insert(bar.FmtBtn, 0)
	bar.Insert(sep, 1)
	bar.Insert(bar.RunBtn, 2)
	return &bar
}
