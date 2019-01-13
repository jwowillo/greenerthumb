all: build/message.a

build/message.a: build/message.o | build
	ar rvs build/message.a build/*.o

build/message.o: message.cc | build
	g++ -c -O2 -std=c++1z message.cc -o build/message.o

clean:
	rm -rf build

build:
	mkdir -p build

