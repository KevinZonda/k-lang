#!/bin/bash
cd ./src
go mod tidy
go build -v -ldflags "-s -w" -o ./out/interpreter ./main/*.go
cd ..