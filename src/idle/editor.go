package idle

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/fmtr"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/gotk3/gotk3/gtk"
	sourceview "github.com/linuxerwang/sourceview3"
	"log"
	"strconv"
	"strings"
)

type EditorW struct {
	*gtk.Window
	CodeE     *CodeEditor
	MenuBar   *gtk.MenuBar
	Toolbar   *ToolBar
	CodeView  *gtk.ScrolledWindow
	VBox      *gtk.Box
	ReplE     *CodeEditor
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
	e.CodeE.buf.SetText(s)
	e.SetTitle("IDLE - " + e.Path)
}

func (e *EditorW) syncCursorPos() {
	buf := e.CodeE.buf
	iter := buf.GetIterAtMark(buf.GetInsert())
	line := iter.GetLine()
	col := iter.GetLineOffset()
	e.StatusBar.Push(0, fmt.Sprintf("Line %d, Col %d", line+1, col+1))
}

func NewEditorW() *EditorW {
	w := EditorW{}
	w.Window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.SetTitle("IDLE")

	w.VBox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	w.CodeView, _ = gtk.ScrolledWindowNew(nil, nil)
	w.CodeView.SetVExpand(true)
	w.CodeE = NewCodeEditor("cpp")
	w.CodeView.Add(w.CodeE)

	w.Toolbar = w.NewToolBar()
	w.MenuBar = NewMenuBar()

	w.StatusBar, _ = gtk.StatusbarNew()
	w.StatusBar.SetBorderWidth(0)
	w.StatusBar.SetMarginBottom(0)
	w.StatusBar.SetMarginTop(0)
	w.StatusBar.SetMarginStart(0)
	w.StatusBar.SetMarginEnd(0)

	w.VBox.Add(w.MenuBar)
	w.VBox.Add(w.Toolbar)

	pane1, _ := gtk.PanedNew(gtk.ORIENTATION_VERTICAL)
	w.ReplE = NewCodeEditor("markdown")
	replW, _ := gtk.ScrolledWindowNew(nil, nil)
	replW.Add(w.ReplE)
	pane1.Pack1(w.CodeView, true, false)
	pane1.Pack2(replW, true, false)

	w.VBox.PackStart(pane1, true, true, 0)
	//w.VBox.PackStart(w.CodeView, true, true, 0)
	w.VBox.Add(w.StatusBar)
	w.Add(w.VBox)
	w.SetDefaultSize(800, 600)

	w.syncCursorPos()
	w.CodeE.SourceView.TextView.Connect("move-cursor", w.syncCursorPos)
	w.CodeE.buf.Connect("end-user-action", w.syncCursorPos)

	w.Toolbar.RunBtn.Connect("clicked", func() {
		w.ReplE.AppendEnd("\n===================NEW RUN===================\n")
		ast, errs := parserHelper.Ast(w.CodeE.Text())
		if len(errs) > 0 {
			sb := strings.Builder{}
			sb.WriteString("Parse failed:\n")
			for idx, err := range errs {
				sb.WriteString("[" + strconv.Itoa(idx) + "] ")
				sb.WriteString(err.Error())
				sb.WriteString("\n")
			}
			w.ReplE.AppendEnd(sb.String())
			//dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, strings.TrimSpace(sb.String()))
			//dialog.SetTitle("Parse Failed")
			//dialog.Run()
			//dialog.Destroy()
			return
		}
		e := eval.New(ast, "")
		isPanic, stdio, panicMsg := runCode(e)
		e = nil
		if !isPanic {
			//dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, stdio)
			//dialog.SetTitle("Result")
			//dialog.Run()
			//dialog.Destroy()
			w.ReplE.AppendEnd(stdio)
		} else {
			msg := stdio + "\nBut with following panic:\n" + panicMsg
			dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, msg)
			dialog.SetTitle("Result with Panic")
			dialog.Run()
			dialog.Destroy()
		}

	})
	w.Toolbar.FmtBtn.Connect("clicked", w.FormatCode)

	return &w
}

func (e *EditorW) FormatCode() {
	code := e.CodeE.Text()
	e.CodeE.buf.SetText(fmtr.Fmt(code))
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

type CodeEditor struct {
	*sourceview.SourceView
	LanguageManager *sourceview.SourceLanguageManager
	buf             *sourceview.SourceBuffer
}

func (ce *CodeEditor) AppendEnd(s string) {
	_, end := ce.buf.GetBounds()
	ce.buf.Insert(end, s)
}
func (ce *CodeEditor) Text() string {
	start, end := ce.buf.GetBounds()
	txt, err := ce.buf.GetText(start, end, false)
	if err != nil {
		log.Println("CodeEditor.Text() failed:", err)
	}
	return txt
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
	ce.buf, _ = sv.GetBuffer()
	ce.buf.SetLanguage(l)
	return &ce
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
