all: build/render.a

build/render.a: build/lodepng.o \
	build/color.o \
	build/gl.o \
	build/primitive.o \
	build/renderer.o | build
	ar rvs build/render.a build/*.o

build/%.o: %.cc | build
	g++ -c -O2 -std=c++1z $< -o $@

build:
	mkdir -p build
