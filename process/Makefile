.PHONY: select flatten summarize filter clean test

all: process test

process: select clean_program summarize filter flatten

select: | build
	$(MAKE) -C select select
	cp -rf select/build/. build/select

# Called clean_program because clean is already a target.
clean_program: | build
	$(MAKE) -C clean clean_program
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

test: test_select test_clean test_filter test_summarize test_flatten

clean:
	rm -rf build
	$(MAKE) -C select clean
	$(MAKE) -C clean clean
	$(MAKE) -C flatten clean
	$(MAKE) -C summarize clean
	$(MAKE) -C filter clean

build:
	mkdir -p build

test_clean: clean_program | build
	$(MAKE) -C clean test
	mkdir -p build/clean
	cp -rf clean/build/. build/clean
	cp test build/test

test_%: % | build
	$(MAKE) -C $* test
	mkdir -p build/$*
	cp -rf $*/build/. build/$*
	cp test build/test
