package idle

import (
	"bytes"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/eval"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/lib/async"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/parserHelper"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (w *EditorW) RunCode() {
	go w.runCode(w.CodeE.Text(), true, "===================NEW RUN===================")
}

func (w *EditorW) RerunCode() {
	w.Stop()
	go w.runCode(w.CodeE.Text(), false, "===================NEW RUN===================")
}

func (w *EditorW) runCode(code string, loadCtx bool, beginMsg string) (rst eval.DetailedRunResult) {
	w.evalIn = &FakeStdIn{buf: &bytes.Buffer{}}
	if w.isRunning {
		return
	}

	defer func() {
		w.evalIn = nil
		w.setRunning(false)
	}()

	w.ReplE.SmartNewLine()
	if beginMsg != "" {
		w.ReplE.AppendEnd(beginMsg + "\n")
	}

	ast, errs := parserHelper.Ast(code)
	if len(errs) > 0 {
		w.ReplE.AppendTag(w.ReplE.Tags["red"], errs.String())
		return
	}
	if !loadCtx || w.e == nil {
		w.e = eval.New()
	}
	ev := w.e
	stdout := w.ReplE.WriterPipe(w.gtkIO)
	ev.SetStdIn(w.evalIn)
	ev.SetStdOut(stdout)
	ev.SetStdErr(stdout)

	ch := make(chan eval.DetailedRunResult, 1)

	cancel := async.AsyncFunc(func() {
		ch <- ev.DoSafely(ast)
	})
	w.cancelFunc = func() {
		w.gtkIO.Lock()
		defer w.gtkIO.Unlock()
		cancel()
	}
	w.setRunning(true)
	rst = <-ch
	close(ch)
	glib.IdleAdd(func() {
		w.ReplE.ScrollToEndUnsafe()
		if rst.IsPanic {
			if w.PanicWithDlg {
				msg := "Code Panicked:\n" + rst.PanicMsg
				dialog := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, msg)
				dialog.SetTitle("Result with Panic")
				dialog.Run()
				dialog.Destroy()
			} else {
				w.ReplE.AppendTagUnsafe(w.ReplE.Tags["red"], "[PANIC RECEIVED] "+rst.PanicMsg+"\n")
				w.ReplE.ScrollToEndUnsafe()
			}
		}
		if beginMsg != "" { // User Invoke Will Handle
			w.startPromptUnsafe()
		}
	})
	return
}
