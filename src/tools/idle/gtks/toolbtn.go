package gtks

import "github.com/gotk3/gotk3/gtk"

func ToolBtnWithIcon(label, icon string, clicked any) *gtk.ToolButton {
	btn, _ := gtk.ToolButtonNew(nil, label)
	btn.SetIconName(icon)
	if clicked != nil {
		btn.Connect("clicked", clicked)
	}
	return btn
}
