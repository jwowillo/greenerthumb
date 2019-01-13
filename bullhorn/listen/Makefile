.PHONY: client server

all: client server

client: | build
	cd client && go build -o ../build/client

server: | build
	cd server && go build -o ../build/server

clean:
	rm -rf build

build:
	mkdir -p build
