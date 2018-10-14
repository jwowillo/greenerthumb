.PHONY: fan test doc build

all: fan test doc

fan: build
	go build -o build/fan

test: fan build
	cp -rf data build
	cp test build/test

doc:
	$(MAKE) -C doc

build:
	mkdir -p build
