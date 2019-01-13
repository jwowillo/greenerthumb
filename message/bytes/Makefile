.PHONY: test

all: bytes test

bytes: build/bytes

build/bytes: *.go ../*.go ../../*.go
	go build -o build/bytes

test: bytes | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

clean:
	rm -rf build

build:
	mkdir -p build
