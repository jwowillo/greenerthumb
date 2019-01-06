.PHONY: test

all: header test

header: build/header

build/header: *.go ../../*.go
	go build -o build/header

test: header | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
