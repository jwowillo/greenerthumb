# `plot`

`plot`s `greenerthumb` JSON messages from STDIN.

## Documentation

`make` in the `doc` directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make plot` builds `plot`.
* `make test` copies a sample log into the build directory with `plot`.

## Running

```
./plot

< {"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
< {"Name": "Air", "Timestamp": 3600, "Temperature": 84.5}
< {"Name": "Soil", "Timestamp": 3600, "Moisture": 0.35}
< {"Name": "Air", "Timestamp": 3600, "Temperature": 82.1}
```

This will plot 2 lines labelled 'Soil Moisture' and 'Air Temperature'. Each will
have 2 points. The 'Soil'-line will start at hour 0 and finish at hour 1. The
'Air'-line will start at hour 1 and finish at hour 2. The entire plot will
occupy 2 hours. The range for 'Soil Moisture' will be [0.35, 0.37] and the range
for 'Air Temperature' will be [82.1, 84.5].

## Dependencies

* OpenGL, GLEW, GLUT, and SOIL are required.
* The JSON parser is SimpleJSON taken from https://github.com/nbsdx/SimpleJSON.

## `gen_icon.py`

`gen_icon.py` is a utility script that generates the icon into the correct c++
file. It should be run whenever the icon is updated.
