package gtks

import "github.com/gotk3/gotk3/gtk"

func NewMenuItem(label string, clicked any) *gtk.MenuItem {
	item, _ := gtk.MenuItemNewWithLabel(label)
	if clicked != nil {
		item.Connect("activate", clicked)
	}
	return item
}

type Menu gtk.Menu

func NewMenu() *Menu {
	menu, _ := gtk.MenuNew()
	return (*Menu)(menu)
}

func (m *Menu) AppendNewMenuItem(label string, clicked any) *Menu {
	i := NewMenuItem(label, clicked)
	m.Append(i)
	return m
}

func (m *Menu) AppendSeparator() *Menu {
	i, _ := gtk.SeparatorMenuItemNew()
	m.Append(i)
	return m
}
