# `sense` Design

`sense` writes `greenerthumb` JSON messages from sensors to STDOUT.

Anomalous readings shouldn't be printed.

Emulated data is printed if sensors aren't detected.

## Program

### `air`

`air` senses the 'Air Status Message' at 0.1 hertz.

![Air Schematic](air.png)

### `soil`

`soil` senses the 'Soil Status Message' at 0.1 hertz.

![Soil Schematic](soil.png)

## Example

### `air`

```
./air

{"Name": "Air", "Timestamp": 0, "Temperature": 67.4}
```

### `soil`

```
./soil

{"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
```
