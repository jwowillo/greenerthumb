.PHONY: build test

all: process summarize filter test_process test_summarize test_filter

process: build
	cd flatten && make
	cp -rf flatten/build/. build/flatten

summarize: build
	cd summarize && make
	cp -rf summarize/build/. build/summarize

filter: build
	cd filter && make
	cp -rf filter/build/. build/filter

test: test_process test_summarize

test_process: build
	cd flatten && make test
	cp -rf flatten/build/. build/flatten
	cp test build/test

test_summarize: build
	cd summarize && make test
	cp -rf summarize/build/. build/summarize
	cp test build/test

test_filter: build
	cd filter && make test
	cp -rf filter/build/. build/filter
	cp test build/test

build:
	mkdir -p build/flatten
	mkdir -p build/summarize
