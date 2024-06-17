# K Language

> [!NOTE]  
> This was also my FYP (Final Year Project, aka thesis) of my undergraduate at the University of Birmingham.  
> The latest version (commit hash) I submitted to the University is `97f26d656a9ff409bae506240da118cb8b10f53f`.
> I also tagged it as [fyp](https://github.com/KevinZonda/k-lang/tree/fyp).  
> You can also access it at uni's GitLab (if you have permission) at <https://git.cs.bham.ac.uk/projects-2023-24/xxs166>

> Instruction to run the code, [click here](Instruction.md)

## Repository Structure

```tree
├───antlr4          # Language's Antlr4 definition is here
├───scripts         # All scripts stores here
├───res             # Resource files
├───src             # Source code of the language
│   ├───ast         # AST Constractor
│   │   ├───node
│   │   ├───...
│   │   └───visitor
│   ├───eval        # Evaluator
│   ├───main
│   ├───out         # Ignored, compiled binary
│   ├───parser      # Antlr4 generated codes
│   ├───...
│   └───utils
├───playground      # Playground: K-Language Web Demo
│   ├───index.html  # Entry Point
│   ├───syntax.js   # Syntax Highlighter
│   └───core.wasm   # K-Language Wasm Core
├───web             # Landing Pages & others
└───vscode-plugin   # VSCode Plugin for this Language
    ├───klang-highlighter
    └───out         # Ignored, but is for generated VSCode plugin file (.VSIX)
```

## Landing Page (Web)

```
cd web
pnpm i
pnpm dev
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

```bash
sudo apt-get install libgtk-3-dev
sudo apt-get install libgtksourceview-3.0-dev
sudo apt-get install pkg-config
sudo apt-get install gcc
```

### Fedora

```bash
sudo dnf install gtksourceview3-devel
sudo dnf install gtk3-devel
sudo dnf install pkgconf
sudo dnf install gcc
sudo ln -s /usr/lib64/pkgconfig/gtksourceview-3.0.pc /usr/lib64/pkgconfig/gtksourceview-3.pc # other distros use 3 instead of 3.0
```
