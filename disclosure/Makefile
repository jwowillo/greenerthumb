.PHONY: test

all: disclosure test

disclosure: build/disclosure

build/disclosure: *.go
	go build -o build/disclosure

test: disclosure | build
	cp test build/test

clean:
	rm -rf build

build:
	mkdir -p build
