.PHONY: test

all: select test

select: build/select

build/select: select.go
	go build -o build/select

test: select | build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
