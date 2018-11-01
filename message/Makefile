.PHONY: build

all: message test

message: build
	cd bytes && go build -o ../build/bytes
	cd json && go build -o ../build/json

test: message build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
