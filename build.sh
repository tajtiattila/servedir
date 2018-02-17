#!/bin/bash

if [[ ! -d release ]]; then
	mkdir release
fi

if [[ ! "$(type -P zip)" ]]; then
	if [[ "$(type -P 7z)" ]]; then
		zip() {
			7z a $@
		}
	else
		echo "no zip"
		exit 1
	fi
fi

GOOS=windows GOARCH=386 go build
zip release/servedir-win32.zip servedir.exe
rm servedir.exe

GOOS=windows GOARCH=amd64 go build
zip release/servedir-win64.zip servedir.exe
rm servedir.exe

