.PHONY: bytes json test

all: message test

message: bytes json

bytes: build
	cd bytes && go build -o ../build/bytes

json: build
	cd json && go build -o ../build/json

test: bytes json build
	rm -rf build/data
	cp -rf data build
	cp test build/test

build:
	mkdir -p build
