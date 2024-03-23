package idle_test

import (
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/tools/idle"
	"github.com/gotk3/gotk3/gtk"
	"os"
	"testing"
)

func TestGui(t *testing.T) {
	os.Setenv("K_LANG_GO_TEST", "1")
	//runtime.LockOSThread()
	gtk.Init(nil)
	w := idle.NewEditorW()
	w.CodeE.WriterPipe().Write([]byte("x = 18"))
	w.RunCode()
	w.RerunCode()
	idle.HelpW().Close()
	w.Close()
	gtk.MainQuit()
	//runtime.UnlockOSThread()

}
