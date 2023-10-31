@echo off
antlr4 -Dlanguage=Go -visitor .\antlr4\V2Lexer.g4 .\antlr4\V2Parser.g4 
rmdir /s /q .\src\parser\
mkdir .\src\parser\
for /r ".\antlr4" %%x in (*.go) do move "%%x" ".\src\parser\"
for /r ".\antlr4" %%x in (*.interp) do move "%%x" ".\src\parser\"
for /r ".\antlr4" %%x in (*.tokens) do move "%%x" ".\src\parser\"