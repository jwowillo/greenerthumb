.PHONY: publish subscribe test

all: bullhorn test

bullhorn: publish subscribe

publish: build
	cd publish && go build -o ../build/publish

subscribe: build
	cd subscribe && go build -o ../build/subscribe

test: publish subscribe build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
