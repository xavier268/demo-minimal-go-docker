#!/bin/bash
echo "Building a windows executable"
GOOS=windows
GOARCH=amd64
#GOARCH=386
go build -o bin/main.exe  ./main/
