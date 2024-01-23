@echo off

cd src
go mod tidy
go build -v -o out/interpreter.exe main/main.go
cd ..