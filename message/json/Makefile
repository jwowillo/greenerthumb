.PHONY: test

all: json test

json: build/json

build/json: *.go ../*.go ../../*.go
	go build -o build/json

test: json | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

clean:
	rm -rf build

build:
	mkdir -p build
