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
