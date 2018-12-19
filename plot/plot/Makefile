all: build/plot.a

build/plot.a: build/plot.o | build
	ar rvs build/plot.a build/*.o

build/plot.o: plot.cc | build
	g++ -c -O2 -std=c++1z plot.cc -o build/plot.o

build:
	mkdir -p build
