.PHONY: clean flatten summarize filter test

all: process test

process: clean summarize filter flatten

clean: | build
	$(MAKE) -C clean clean
	cp -rf clean/build/. build/clean

flatten: | build
	$(MAKE) -C flatten flatten
	cp -rf flatten/build/. build/flatten

summarize: | build
	$(MAKE) -C summarize summarize
	cp -rf summarize/build/. build/summarize

filter: | build
	$(MAKE) -C filter filter
	cp -rf filter/build/. build/filter

test: test_clean test_filter test_summarize test_flatten

build:
	mkdir -p build

test_%: % | build
	$(MAKE) -C $* test
	cp -rf $*/build/. build/$*
	cp test build/test
