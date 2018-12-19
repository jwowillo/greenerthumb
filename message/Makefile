.PHONY: bytes json test

all: message test

message: bytes json

bytes: build
	$(MAKE) -C bytes
	cp -rf bytes/build/. build

json: build
	$(MAKE) -C json
	cp -rf json/build/. build

test: bytes json build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
