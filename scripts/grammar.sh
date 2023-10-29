#!/bin/bash
antlr -Dlanguage=Go ./antlr4/V2Lexer.g4 ./antlr4/V2Parser.g4 
rm -fr ./klang/parser
mkdir ./klang/parser
mv ./antlr4/*.go     ./klang/parser/
mv ./antlr4/*.interp ./klang/parser/
mv ./antlr4/*.tokens ./klang/parser/
