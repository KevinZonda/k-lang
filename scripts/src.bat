@echo off

cd src
go mod tidy
go build -v -o out/interpreter.exe main/cli/main.go
cd ..