//go:build !darwin

package idle

type Customizer struct {
}

func (w *EditorW) onNewWindow() bool {
	return true
}

func (w *EditorW) onWindowCreated() {
}
