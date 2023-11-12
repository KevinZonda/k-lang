#!/bin/bash
antlr -Dlanguage=Go -visitor ./antlr4/V2Lexer.g4 ./antlr4/V2Parser.g4 
rm -fr ./src/parser
mkdir ./src/parser
mv ./antlr4/*.go     ./src/parser/
mv ./antlr4/*.interp ./src/parser/
mv ./antlr4/*.tokens ./src/parser/
