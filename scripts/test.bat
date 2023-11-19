@echo off
cd ./src
go test -v -coverprofile cover.out ./...
go tool cover -html cover.out -o cover.html
del cover.out
set f=file:///%cd%/cover.html
start chrome %f%
cd ..