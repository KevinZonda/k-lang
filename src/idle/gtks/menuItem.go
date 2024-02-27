package gtks

import "github.com/gotk3/gotk3/gtk"

func NewMenuItem(label string, clicked any) *gtk.MenuItem {
	item, _ := gtk.MenuItemNewWithLabel(label)
	if clicked != nil {
		item.Connect("activate", clicked)
	}
	return item
}
