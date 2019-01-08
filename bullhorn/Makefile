.PHONY: broadcast listen pubsub test

all: bullhorn test

bullhorn: broadcast listen pubsub

broadcast: | build
	$(MAKE) -C broadcast
	cp -rf broadcast/build/. build/broadcast

listen: | build
	$(MAKE) -C listen
	cp -rf listen/build/. build/listen

pubsub: | build
	$(MAKE) -C pubsub
	cp -rf pubsub/build/. build/pubsub

test: bullhorn | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
