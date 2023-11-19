#!/bin/bash
cd ./src
go test -v -coverprofile cover.out ./...
go tool cover -html cover.out -o cover.html
rm cover.out
start chrome "file:///$PWD/cover.html"
cd ..