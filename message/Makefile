.PHONY: header bytes json test

all: message test

message: header bytes json

header: | build
	$(MAKE) -C header
	cp -rf header/build/. build/header

bytes: | build
	$(MAKE) -C bytes
	cp -rf bytes/build/. build/bytes

json: | build
	$(MAKE) -C json
	cp -rf json/build/. build/json

test: test_header test_bytes test_json

clean:
	rm -rf build
	$(MAKE) -C header clean
	$(MAKE) -C bytes clean
	$(MAKE) -C json clean

build:
	mkdir -p build

test_%: % | build
	$(MAKE) -C $* test
	cp -rf $*/build/. build/$*
	cp test build/test
