.PHONY: publish subscribe yell snoop test

all: bullhorn test

bullhorn: pubsub broadcast

pubsub: publish subscribe

broadcast: yell snoop

publish: | build
	cd publish && go build -o ../build/publish

subscribe: | build
	cd subscribe && go build -o ../build/subscribe

yell: | build
	cd yell && go build -o ../build/yell

snoop: | build
	cd snoop && go build -o ../build/snoop

test: publish subscribe yell build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
