#!/bin/bash
cd ./src
go build -v -o ./out/interpreter ./main/*.go
cd ..