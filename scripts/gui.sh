#!/bin/bash

cd src
go mod tidy
go build -v -ldflags "-H windowsgui" -o out/idle.exe main/gui/main.go
cd ..