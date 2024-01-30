package gtks

import (
	sourceview "github.com/linuxerwang/sourceview3"
	"log"
	"strings"
)

type CodeEditor struct {
	*sourceview.SourceView
	LanguageManager *sourceview.SourceLanguageManager
	Buf             *sourceview.SourceBuffer
}

func (ce *CodeEditor) ScrollToEnd() {
	_, end := ce.Buf.GetBounds()
	ce.ScrollToIter(end, 0, false, 0, 0)
}

func (ce *CodeEditor) SmartNewLine() {
	t := ce.Text()
	if !strings.HasSuffix(t, "\n") && len(t) > 0 {
		ce.AppendEnd("\n")
	}
}

func (ce *CodeEditor) AppendEnd(s string) {
	_, end := ce.Buf.GetBounds()
	ce.Buf.Insert(end, s)
}
func (ce *CodeEditor) Text() string {
	start, end := ce.Buf.GetBounds()
	txt, err := ce.Buf.GetText(start, end, false)
	if err != nil {
		log.Println("CodeEditor.Text() failed:", err)
	}
	return txt
}

func (ce *CodeEditor) SetText(s string) {
	ce.Buf.SetText(s)
}

func NewCodeEditor(lang string) *CodeEditor {
	sv, _ := sourceview.SourceViewNew()
	lm, _ := sourceview.SourceLanguageManagerGetDefault()
	ce := CodeEditor{
		SourceView:      sv,
		LanguageManager: lm,
	}

	SetCss(ce, `
#CodeEditor {
	font-size: 14px;
	font-family: monospace;
}
`)

	ce.SetName("CodeEditor")
	l, _ := lm.GetLanguage(lang)
	sv.SetShowLineNumbers(true)
	ce.Buf, _ = sv.GetBuffer()
	ce.Buf.SetLanguage(l)
	return &ce
}
