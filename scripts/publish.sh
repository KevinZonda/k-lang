#!/bin/bash

cd ./src
go mod tidy

rm -fr ./out/publish
mkdir -p ./out/publish

hash=$(git log --pretty=format:'%h' -n 1)
date=$(date +"%d/%m/%Y")
time=$(date +"%T")
branch=$(git branch --show-current)

flag="-X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.CommitHash=$hash"
flag=" $flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildDate=$date"
flag=" $flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildTime=$time"
flag=" $flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildBranch=$branch"

flag=" $flag -s -w"

echo 'Publishing macOS...'
GOOS=darwin GOARCH=arm64 go build -v -ldflags "$flag" -o ./out/publish/k-std-darwin-arm64 ./main/cli/*.go
GOOS=darwin GOARCH=amd64 go build -v -ldflags "$flag" -o ./out/publish/k-std-darwin-amd64 ./main/cli/*.go

echo 'Publishing Windows...'
GOOS=windows GOARCH=amd64 go build -v -ldflags "$flag" -o ./out/publish/k-std-windows-arm64.exe ./main/cli/*.go
GOOS=windows GOARCH=amd64 go build -v -ldflags "$flag" -o ./out/publish/k-std-windows-amd64.exe ./main/cli/*.go
GOOS=windows GOARCH=386   go build -v -ldflags "$flag" -o ./out/publish/k-std-windows-386.exe ./main/cli/*.go

echo 'Publishing Linux...'
GOOS=linux GOARCH=amd64 go build -v -ldflags "$flag" -o ./out/publish/k-std-linux-arm64 ./main/cli/*.go
GOOS=linux GOARCH=amd64 go build -v -ldflags "$flag" -o ./out/publish/k-std-linux-amd64 ./main/cli/*.go
GOOS=linux GOARCH=386   go build -v -ldflags "$flag" -o ./out/publish/k-std-linux-386 ./main/cli/*.go

cd ..