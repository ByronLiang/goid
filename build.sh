#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/$1/main.go

rm -rf bin/$1
mkdir -p bin/$1
mv main bin/$1
cp Dockerfile bin/$1
