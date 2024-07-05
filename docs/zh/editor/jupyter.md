# Integration with Jupyter Lab/Notebooks

## Install Jupyter Lab

```bash
pip install jupyter
pip install jupyterlab
```

## Install Jupyter Kernel

### Windows User

```cmd
mkdir %APPDATA%\jupyter\kernels\k
xcopy %SRC/scripts/jupyter/ %APPDATA%\jupyter\kernels\k /s
```

Modify the `%APPDATA%\jupyter\kernels\k\kernel.json`:

```json
{
    "argv": [
        "<YOUR KLANG STANDARD BINARY FILE PATH>",
        "jupyter",
        "{connection_file}"
    ],
    "display_name": "K Kernel",
    "language": "k",
    "name": "k"
}
```

and `%APPDATA%\jupyter\kernels\k\kernel.json.in`

```json
{
    "argv": [
        "<YOUR KLANG STANDARD BINARY FILE PATH>",
        "jupyter"
    ],
    "display_name": "K Kernel",
    "language": "k",
    "name": "k"
}
```

### macOS User

```sh
mkdir -p ~/Library/Jupyter/kernels/k
cp  $SRC/scripts/jupyter/* ~/Library/Jupyter/kernels/k
```

Modify the `~/Library/Jupyter/kernels/k/kernel.json`:

```json
{
    "argv": [
        "<YOUR KLANG STANDARD BINARY FILE PATH>",
        "jupyter",
        "{connection_file}"
    ],
    "display_name": "K Kernel",
    "language": "k",
    "name": "k"
}
```

and `~/Library/Jupyter/kernels/k/kernel.json.in`

```json
{
    "argv": [
        "<YOUR KLANG STANDARD BINARY FILE PATH>",
        "jupyter"
    ],
    "display_name": "K Kernel",
    "language": "k",
    "name": "k"
}
```

### Linux User

```
mkdir -p ~/.local/share/jupyter/kernels/k
cp  $SRC/scripts/jupyter/* ~/.local/share/jupyter/kernels/k
```

Modify the `~/.local/share/jupyter/kernels/k/kernel.json`:

```json
{
    "argv": [
        "<YOUR KLANG STANDARD BINARY FILE PATH>",
        "jupyter",
        "{connection_file}"
    ],
    "display_name": "K Kernel",
    "language": "k",
    "name": "k"
}
```

and `~/.local/share/jupyter/kernels/k/kernel.json.in`

```json
{
    "argv": [
        "<YOUR KLANG STANDARD BINARY FILE PATH>",
        "jupyter"
    ],
    "display_name": "K Kernel",
    "language": "k",
    "name": "k"
}
```