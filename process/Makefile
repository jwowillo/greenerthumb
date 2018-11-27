.PHONY: build

all: process test

process: build
	cd flatten && make
	cp -rf flatten/build/. build/flatten

test: build
	cd flatten && make test
	cp -rf flatten/build/. build/flatten
	cp test build/test

build:
	mkdir -p build/flatten
