#!/bin/bash

cd src
go mod tidy

hash=$(git log --pretty=format:'%h' -n 1)
date=$(date +"%d/%m/%Y-%T")
branch=$(git branch --show-current)
flag="-X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.CommitHash=$hash"
flag=" $flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildTime=$date"
flag=" $flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildBranch=$branch"
#flag=" $flag -H windowsgui"

go build -v -ldflags "$flag" -o out/idle main/gui/main.go
cd ..