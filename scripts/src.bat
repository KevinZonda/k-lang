@echo off

cd src
go build -v -o out/interpreter.exe main/main.go
cd ..