.PHONY: read write test

all: store test

store: read write

read: | build
	$(MAKE) -C read
	cp -rf read/build/. build

write: | build
	$(MAKE) -C write
	cp -rf write/build/. build

test: store | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
