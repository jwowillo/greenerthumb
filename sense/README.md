# `sense`

`sense` programs write `greenerthumb` JSON messages from sensors to STDOUT.

# Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make sense` builds all sense programs.

## Running

### Air

```
./air ?--rate <rate> ?--pin <pin>
```

An example is:

```
./air

{"Name": "Air", "Timestamp": 0, "Temperature": 67.4}
```

### Soil

```
./soil ?--rate <rate> \
    ?--channel <channel> \
    ?--clk <clk> \
    ?--miso <miso> \
    ?--mosi <mosi> \
    ?--cs <cs>
```

An example is:

```
./soil

{"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
```

Each accepts all necessary GPIO pin numbers, ADC channels, and rates as optional
flags with defaults chosen from the schematics.

## Emulators

Emulators are provided for all programs and each accepts an optional rate flag.

```
./<sensor> ?--rate <rate>
```
