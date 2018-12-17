.PHONY: plot

all: plot test

plot: | build
	$(MAKE) -C message
	cp message/build/message.a build
	$(MAKE) -C plot
	cp plot/build/plot.a build
	$(MAKE) -C render
	cp render/build/render.a build
	g++ -O2 -std=c++1z main.cc build/*.a -o build/plot \
		-lGL \
		-lglut \
		-lGLU \
		-lglfw \
		-l GLEW \
		-lX11 \
		-lXxf86vm \
		-lXrandr \
		-lpthread \
		-lXi

test: plot | build
	rm -rf build/data
	cp -rf data build

build:
	mkdir -p build
