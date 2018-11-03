.PHONY: build

all: bullhorn test

bullhorn: build
	cd publish && go build -o ../build/publish
	cd subscribe && go build -o ../build/subscribe

test: bullhorn build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
