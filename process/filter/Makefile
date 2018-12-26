.PHONY: test

all: filter test

filter: build/filter

build/filter: filter.go
	go build -o build/filter

test: filter | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
