.PHONY: test build

all: summarize test

summarize: build
	go build -o build/summarize

test: summarize build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
