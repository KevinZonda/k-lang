@echo off

cd src
go mod tidy
go build -v -ldflags "-s -w" -o out/interpreter.exe main/main.go
cd ..