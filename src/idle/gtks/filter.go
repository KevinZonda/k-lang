package gtks

import "github.com/gotk3/gotk3/gtk"

func NewFileFilter(pattern, description string) *gtk.FileFilter {
	filter, _ := gtk.FileFilterNew()
	filter.AddPattern(pattern)
	filter.SetName(description)
	return filter
}
