# Playground of K Language



Copy`"wasm_exec.js` from `$GOROOT/misc/wasm` (in Windows, default is in `C:\Program Files\Go\misc\wasm`) to this folder.

Build a wasm called `interpreter.wasm` by following command in `src` folder

```pwsh
$Env:GOOS = "js"; $Env:GOARCH = "wasm"
go build -v -o ../playground/interpreter.wasm main/wasm/main_wasm.go
```

In Linux/Unix/macOS, use

```
GOOS=js GOARCH=wasm go build -v -o ../playground/interpreter.wasm main/wasm/main_wasm.go
```

