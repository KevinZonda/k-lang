package gtks

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
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
	glib.IdleAdd(func() {
		_, end := ce.Buf.GetBounds()
		ce.ScrollToIter(end, 0, false, 0, 0)
	})
}

func (ce *CodeEditor) SmartNewLine() {
	glib.IdleAdd(func() {
		t := ce.Text()
		if !strings.HasSuffix(t, "\n") && len(t) > 0 {
			_, end := ce.Buf.GetBounds()
			ce.Buf.Insert(end, "\n")
		}
	})

}

func (ce *CodeEditor) SmartNewLineUnsafe() {
	t := ce.Text()
	if !strings.HasSuffix(t, "\n") && len(t) > 0 {
		_, end := ce.Buf.GetBounds()
		ce.Buf.Insert(end, "\n")
	}

}

func (ce *CodeEditor) AppendEnd(s string) {
	glib.IdleAdd(func() {
		_, end := ce.Buf.GetBounds()
		ce.Buf.Insert(end, s)
	})
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

func (ce *CodeEditor) AppendTag(tag *gtk.TextTag, content string) {
	glib.IdleAdd(func() {
		ce.AppendTagUnsafe(tag, content)
	})
}

func (ce *CodeEditor) AppendTagUnsafe(tag *gtk.TextTag, content string) {
	glib.IdleAdd(func() {
		ce.Buf.InsertWithTag(ce.Buf.GetEndIter(), content, tag)
	})
}

func NewCodeEditor(lang string, fontSize int) *CodeEditor {
	if fontSize <= 0 {
		fontSize = 14
	}
	sv, _ := sourceview.SourceViewNew()
	lm, _ := sourceview.SourceLanguageManagerGetDefault()
	ce := CodeEditor{
		SourceView:      sv,
		LanguageManager: lm,
	}

	SetCss(ce, fmt.Sprintf(`
#CodeEditor {
	font-size: %dpx;
	font-family: monospace;
}
`, fontSize))

	ce.SetName("CodeEditor")
	l, _ := lm.GetLanguage(lang)
	sv.SetShowLineNumbers(true)
	ce.Buf, _ = sv.GetBuffer()
	ce.Buf.SetLanguage(l)
	return &ce
}

func (ce *CodeEditor) NewTextTag(name string, color string) *gtk.TextTag {
	return ce.Buf.CreateTag(name, map[string]any{
		"foreground": color,
	})
}
