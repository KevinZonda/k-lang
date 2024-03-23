package gtks

import "github.com/gotk3/gotk3/gtk"

type ICSSAble interface {
	GetStyleContext() (*gtk.StyleContext, error)
}

func SetCss(w ICSSAble, css string) {
	cssP, _ := gtk.CssProviderNew()
	cssP.LoadFromData(css)
	style, _ := w.GetStyleContext()
	style.AddProvider(cssP, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}
