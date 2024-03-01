package idle

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"runtime"
)

func (w *EditorW) shortcut(acc string, f func()) {
	key, mod := gtk.AcceleratorParse(acc)
	w.acc.Connect(key, mod, gtk.ACCEL_VISIBLE, f)
}

func ctrl() string {
	if runtime.GOOS == "darwin" {
		return "<Primary>"
	}
	return "<Control>"
}

func (w *EditorW) startPromptUnsafe() {
	w.ReplE.SmartNewLineUnsafe()
	w.ReplE.AppendTagUnsafe(w.ReplE.Tags["blue"], ">>> ")
	w.ReplE.ScrollToEndUnsafe()
}

func (w *EditorW) startPrompt() {
	glib.IdleAdd(w.startPromptUnsafe)
}

func (w *EditorW) syncCursorPos() {
	buf := w.CodeE.Buf
	iter := buf.GetIterAtMark(buf.GetInsert())
	line := iter.GetLine()
	col := iter.GetLineOffset()
	w.StatusBar.Push(0, fmt.Sprintf("Line %d, Col %d", line+1, col+1))
}

func (w *EditorW) syncRunningStat() {
	if !w.isRunning {
		w.cancelFunc = nil
	}
	glib.IdleAdd(func() {
		w.Toolbar.RunBtn.SetSensitive(!w.isRunning)
		w.Toolbar.StopBtn.SetSensitive(w.isRunning)
	})
}
