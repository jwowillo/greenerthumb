.PHONY: test build

all: fan test

fan: build
	go build -o build/fan

test: fan build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
