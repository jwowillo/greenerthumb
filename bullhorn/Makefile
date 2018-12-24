.PHONY: publish subscribe yell snoop talk listen test

all: bullhorn test

bullhorn: pubsub broadcast reliable

pubsub: publish subscribe

broadcast: yell snoop

reliable: listen talk

publish: | build
	cd publish && go build -o ../build/publish

subscribe: | build
	cd subscribe && go build -o ../build/subscribe

yell: | build
	cd yell && go build -o ../build/yell

snoop: | build
	cd snoop && go build -o ../build/snoop

listen: | build
	cd listen && go build -o ../build/listen

talk: | build
	cd talk && go build -o ../build/talk

test: bullhorn | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
