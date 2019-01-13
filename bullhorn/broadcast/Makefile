.PHONY: client server

all: client server

client: build/client

build/client: client/*.go ../*.go ../../*.go
	cd client && go build -o ../build/client

server: build/server

build/server: server/*.go ../*.go ../../*.go
	cd server && go build -o ../build/server

clean:
	rm -rf build
