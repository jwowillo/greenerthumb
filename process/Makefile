.PHONY: build test

all: process test

process: clean summarize filter flatten

clean: build
	cd clean && make
	cp -rf clean/build/. build/clean

flatten: build
	cd flatten && make
	cp -rf flatten/build/. build/flatten

summarize: build
	cd summarize && make
	cp -rf summarize/build/. build/summarize

filter: build
	cd filter && make
	cp -rf filter/build/. build/filter

test: test_clean test_filter test_summarize test_flatten

test_clean: build
	cd clean && make test
	cp -rf clean/build/. build/clean
	cp test build/test

test_flatten: build
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
