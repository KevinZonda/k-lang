//go:build darwin

/*
 * This file originally comes from:
 * - https://github.com/coyim/gotk3osx/blob/main/menu_item.go
 * - https://github.com/coyim/gotk3osx/blob/main/menu_shell.go
 * - https://github.com/coyim/gotk3osx/blob/main/pixbuf.go
 * - https://github.com/coyim/gotk3osx/blob/main/widget.go
 *
 * Modified to remove other project specific dependencies
 * Modified by Xiang Shi (xxs166@student.bham.ac.uk)
 */

package gotk3osx

// #include <gtk/gtk.h>
// #include <gdk/gdk.h>
// #include <gtkosxapplication.h>
// #include "gtkosx.go.h"
import "C"
import (
	"github.com/gotk3/gotk3/gdk"
	"unsafe"
)
import "github.com/gotk3/gotk3/gtk"

func nativeMenuShell(v *gtk.MenuShell) *C.GtkMenuShell {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuShell(p)
}

func nativeMenuItem(v *gtk.MenuItem) *C.GtkMenuItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenuItem(p)
}

func nativePixbuf(v *gdk.Pixbuf) *C.GdkPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkPixbuf(p)
}

func nativeWidget(v *gtk.Widget) *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWidget(p)
}

func GtkMenuToGtkMenuShell(v *gtk.Menu) *gtk.MenuShell {
	m := v.MenuShell
	return &m
}
