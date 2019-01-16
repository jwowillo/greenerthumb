all: sense

sense: emulate air-sensor soil-sensor

emulate: build
	cp emulate-air-sensor build/emulate-air-sensor
	cp emulate-soil-sensor build/emulate-soil-sensor

air-sensor: build
	cp air-sensor build/air-sensor

soil-sensor: build
	cp soil-sensor build/soil-sensor

clean:
	rm -rf build

build:
	mkdir -p build
