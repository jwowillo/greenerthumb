.PHONY: test build

all: clean test

clean: build
	go build -o build/clean

test: clean build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
