package gtks

import "github.com/gotk3/gotk3/gtk"

func WrapToScrolledWindow(w gtk.IWidget) *gtk.ScrolledWindow {
	sw, _ := gtk.ScrolledWindowNew(nil, nil)
	sw.Add(w.(gtk.IWidget))
	return sw
}
