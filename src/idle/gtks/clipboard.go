package gtks

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func Clipboard() *gtk.Clipboard {
	c, _ := gtk.ClipboardGet(gdk.SELECTION_CLIPBOARD)
	return c
}
