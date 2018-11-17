.PHONY: build

all: plot test

plot: build
	g++ -O2 -std=c++1z main.cc render/lodepng.cpp -o build/plot \
		-lGL \
		-lglut \
		-lGLU \
		-lglfw \
		-lGLEW \
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
