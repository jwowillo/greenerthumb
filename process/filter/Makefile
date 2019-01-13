.PHONY: test

all: filter test

filter: build/filter

build/filter: *.go ../*.go ../../*.go
	go build -o build/filter

test: filter | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

clean:
	rm -rf build

build:
	mkdir -p build
