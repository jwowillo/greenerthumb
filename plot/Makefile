.PHONY: build

all: plot test

plot: build
	g++ -O2 -std=c++1z main.cc -o build/plot \
		-lGL \
		-lglut \
		-lGLU \
		-lglfw \
		-lGLEW \
		-lSOIL \
		-lX11 \
		-lXxf86vm \
		-lXrandr \
		-lpthread \
		-lXi

test: plot build
	rm -rf build/data
	cp -rf data build

build:
	mkdir -p build
