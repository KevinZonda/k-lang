//go:build !darwin

package idle

type Customizer struct {
}

func (w *EditorW) onNewWindow() {
}

func (w *EditorW) onWindowCreated() {
}
