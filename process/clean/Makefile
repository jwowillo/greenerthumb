.PHONY: test

all: clean test

clean: build/clean

build/clean: calculate.go clean.go
	go build -o build/clean

test: clean | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
