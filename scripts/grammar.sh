#!/bin/bash
antlr -Dlanguage=Go -visitor ./antlr4/V2Lexer.g4 ./antlr4/V2Parser.g4 
rm -fr ./src/parser
mkdir ./src/parser
mv ./src/*.go     ./klang/parser/
mv ./src/*.interp ./klang/parser/
mv ./src/*.tokens ./klang/parser/
