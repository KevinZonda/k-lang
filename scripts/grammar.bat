@echo off
antlr4 -Dlanguage=Go -visitor .\antlr4\V2Lexer.g4 .\antlr4\V2Parser.g4 
rmdir /s /q .\klang\parser\
mkdir .\klang\parser\
for /r ".\antlr4" %%x in (*.go) do move "%%x" ".\klang\parser\"
for /r ".\antlr4" %%x in (*.interp) do move "%%x" ".\klang\parser\"
for /r ".\antlr4" %%x in (*.tokens) do move "%%x" ".\klang\parser\"