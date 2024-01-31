#!/bin/bash

cd src
go mod tidy
go build -v -ldflags "-H windowsgui" -o out/idle.exe idle/main/main.go
cd ..