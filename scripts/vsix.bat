@echo off
cd .\vscode-plugin\klang-highlighter
vsce package
cd ..
cd ..
for /r ".\vscode-plugin\klang-highlighter" %%x in (*.vsix) do move "%%x" ".\vscode-plugin\out\"
