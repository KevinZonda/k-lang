package main

import (
	"errors"
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"os"
)

func CreateDesktopIcon() {
	const Path string = "/usr/share/applications/k-lang-idle.desktop"

	_, err := os.Stat(Path)
	if !errors.Is(err, os.ErrNotExist) {
		return
	}
	iox.WriteAllText(Path, fmt.Sprintf(`[Desktop Entry]
Name=K Language IDLE
Comment=K Language's light-weigh editor
GenericName=Text Editor
Exec=%s
Type=Application
Categories=TextEditor;Development;IDE;
Keywords=vscode;
`, os.Args[0]))
}
