.PHONY: build doc

all: greenerthumb test doc

greenerthumb: build
	$(MAKE) -C fan fan
	cp -rf fan/build build/fan

test: build
	$(MAKE) -C fan test
	cp -rf fan/build build/fan

build:
	mkdir -p build

doc:
	$(MAKE) -C doc
