.PHONY: test

all: flatten test

flatten: build
	go build -o build/flatten

test: flatten build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
