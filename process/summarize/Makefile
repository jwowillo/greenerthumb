.PHONY: test

all: summarize test

summarize: build/summarize

build/summarize: *.go ../*.go ../../*.go
	go build -o build/summarize

test: summarize | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
