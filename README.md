# Final Year Project

## Repository Structure

```tree
├───antlr4          # Language's Antlr4 definition is here
├───scripts         # All scripts stores here
├───src             # Source code of the language
│   ├───ast         # AST Constractor
│   │   ├───node
│   │   ├───...
│   │   └───visitor
│   ├───eval        # Evaluator
│   ├───main
│   ├───out         # Ignored, but is for compiled binary
│   ├───parser      # Antlr4 generated codes
│   ├───...
│   └───utils
└───vscode-plugin   # VSCode Plugin for this Language
    ├───klang-highlighter
    └───out         # Ignored, but is for generated VSCode plugin file (.VSIX)
```

## Compile IDLE

IDLE is a GUI-based application, it uses cgo & GTK3 to create the GUI.  
In different platform, GUI config may different.

### Windows

1. Install MSYS2
2. Open MSYS2 MINGW64
3. ```bash
   pacman -S mingw-w64-x86_64-gcc
   pacman -S mingw-w64-x86_64-gtk3
   pacman -S mingw-w64-x86_64-pkg-config
   pacman -S mingw-w64-x86_64-toolchain
   pacman -S mingw-w64-x86_64-gtksourceview3
   ```
4. Add `C:\msys64\mingw64\bin` to PATH

### macOS

1. Install Homebrew
2. ```bash
   brew install gtk+3
   brew install gtksourceview3
   brew install gtk-mac-integration
   brew install pkg-config
   brew install gcc
   ```
   
### Debian-based Linux (Ubuntu, etc.)

1. Install GTK3 & GTKSourceView3
2. ```bash
   sudo apt-get install libgtk-3-dev
   sudo apt-get install libgtksourceview-3.0-dev
   sudo apt-get install pkg-config
   sudo apt-get install gcc
   ```
