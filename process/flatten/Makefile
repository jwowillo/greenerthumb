.PHONY: test

all: flatten test

flatten: build/flatten

build/flatten: *.go ../*.go ../../*.go
	go build -o build/flatten

test: flatten | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
