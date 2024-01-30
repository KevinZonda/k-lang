package idle

import (
	"fmt"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/gotk3/gotk3/gtk"
	sourceview "github.com/linuxerwang/sourceview3"
	"log"
	"strconv"
	"strings"
)

type EditorW struct {
	*gtk.Window
	CodeE    *CodeEditor
	Toolbar  *ToolBar
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

	w.Toolbar.RunBtn.Connect("clicked", func() {
		ast, errs := parserHelper.Ast(w.CodeE.Text())
		if len(errs) > 0 {
			sb := strings.Builder{}
			sb.WriteString("Parse failed:\n")
			for idx, err := range errs {
				sb.WriteString("[" + strconv.Itoa(idx) + "] ")
				sb.WriteString(err.Error())
				sb.WriteString("\n")
			}
			dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, strings.TrimSpace(sb.String()))
			dialog.SetTitle("Parse Failed")
			dialog.Run()
			dialog.Destroy()
			return
		}
		e := eval.New(ast, "")
		isPanic, stdio, panicMsg := runCode(e)
		e = nil
		if !isPanic {
			dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, stdio)
			dialog.SetTitle("Result")
			dialog.Run()
			dialog.Destroy()
		} else {
			msg := stdio + "\nBut with following panic:\n" + panicMsg
			dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, msg)
			dialog.SetTitle("Result with Panic")
			dialog.Run()
			dialog.Destroy()
		}

	})
	return &w
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
}

func NewToolBar() *ToolBar {
	bar := ToolBar{}
	bar.Toolbar, _ = gtk.ToolbarNew()
	bar.SetStyle(gtk.TOOLBAR_BOTH)

	bar.RunBtn, _ = gtk.ToolButtonNew(nil, "Run")
	bar.RunBtn.SetIconName("media-playback-start")

	sep, _ := gtk.SeparatorToolItemNew()
	bar.Insert(sep, 0)
	bar.Insert(bar.RunBtn, 1)
	return &bar
}
