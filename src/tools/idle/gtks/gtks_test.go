package gtks_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/idle/gtks"
	"github.com/gotk3/gotk3/gtk"
	"testing"
)

func TestGtksNotNil(t *testing.T) {
	gtk.Init(nil)
	if gtks.Clipboard() == nil {
		t.Error("Clipboard() should not be nil")
	}
	if gtks.NewFileFilter("*.k", "K files") == nil {
		t.Error("NewFileFilter() should not be nil")
	}

}
