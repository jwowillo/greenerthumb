.PHONY: plot

all: plot test

plot: libraries | build
	g++ -O2 -std=c++1z main.cc build/*.a -o build/plot \
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

test: libraries | build
	g++ -O2 -std=c++1z test.cc build/*.a -o build/test \
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

libraries:
	$(MAKE) -C message
	cp message/build/message.a build
	$(MAKE) -C plot
	cp plot/build/plot.a build
	$(MAKE) -C render
	cp render/build/render.a build

build:
	mkdir -p build
