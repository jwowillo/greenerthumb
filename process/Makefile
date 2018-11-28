.PHONY: build test

all: process summarize test_process test_summarize

process: build
	cd flatten && make
	cp -rf flatten/build/. build/flatten

summarize:
	cd summarize && make
	cp -rf summarize/build/. build/summarize

test: test_process test_summarize

test_process: build
	cd flatten && make test
	cp -rf flatten/build/. build/flatten
	cp test build/test

test_summarize: build
	cd summarize && make test
	cp -rf summarize/build/. build/summarize
	cp test build/test

build:
	mkdir -p build/flatten
	mkdir -p build/summarize
