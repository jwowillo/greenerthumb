# `sense`

`sense` writes `greenerthumb` JSON messages from sensors to STDOUT.

# Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make sense` builds `air` and `soil`.

## Running

```
./air

{"Name": "Air", "Timestamp": 0, "Temperature": 67.4, "Humidity": 0.58}
```

```
./soil

{"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
```

Each accepts all necessary GPIO pin numbers, ADC channels, and rates as optional
flags with defaults chosen from the schematics.
