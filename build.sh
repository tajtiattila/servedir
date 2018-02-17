#!/bin/bash

if [ -n release ]; then
	mkdir release
fi

GOOS=windows GOARCH=386 go build
zip release/servedir-win32.zip servedir.exe
rm servedir.exe

GOOS=windows GOARCH=amd64 go build
zip release/servedir-win64.zip servedir.exe
rm servedir.exe

