# `device`

`device` programs write `greenerthumb` JSON message bodys from sensors to
STDOUT.

# Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make device` builds all device programs.

## Running

### `air-sensor`

```
./air-sensor ?--rate <rate> ?--pin <pin>
```

An example is:

```
./air-sensor

{"Temperature": 67.4}
```

### `soil-sensor`

```
./soil-sensor ?--rate <rate> \
    ?--channel <channel> \
    ?--clk <clk> \
    ?--miso <miso> \
    ?--mosi <mosi> \
    ?--cs <cs>
```

An example is:

```
./soil-sensor

{"Moisture": 0.37}
```

Each accepts all necessary GPIO pin numbers, ADC channels, and rates as optional
flags with defaults chosen from the schematics.

## Emulators

Emulators are provided for all programs and each accepts an optional rate flag.

```
./<device> ?--rate <rate>
```
