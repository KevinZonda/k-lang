#!/bin/bash
cd ./src
go mod tidy
go build -v -o ./out/interpreter ./main/cli/*.go
cd ..