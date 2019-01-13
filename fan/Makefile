.PHONY: test

all: fan test

fan: build
	go build -o build/fan

test: fan build
	rm -rf build/data
	cp -rf data build
	cp test build/test

clean:
	rm -rf build

build:
	mkdir -p build
