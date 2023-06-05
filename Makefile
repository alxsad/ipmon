.DEFAULT_GOAL := help

prepare:
	rm -rf build
	mkdir build

compile:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -v -trimpath -o build/ipmon-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -v -trimpath -o build/ipmon-darwin-arm64 .

build: prepare compile

local-release:
	goreleaser release --snapshot --clean

remote-release:
	goreleaser release --clean

help:
	cat Makefile
