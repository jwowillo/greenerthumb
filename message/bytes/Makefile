.PHONY: test

all: bytes test

bytes: build/bytes

build/bytes: *.go ../*.go ../../*.go
	go build -o build/bytes

test: bytes | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
