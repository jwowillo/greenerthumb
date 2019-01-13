.PHONY: test

all: clean test

clean_program: build/clean

build/clean: *.go ../*.go ../../*.go
	go build -o build/clean

test: clean_program | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

clean:
	rm -rf build

build:
	mkdir -p build
