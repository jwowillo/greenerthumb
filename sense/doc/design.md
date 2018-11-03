# `sense` Design

`sense` writes `greenerthumb` JSON messages from sensors to STDOUT.

## Program

### `air`

`air` senses the 'Air Status Message' at 1 hert.

![Air Schematic](air.png)

### `soil`

`soil` senses the 'Soil Status Message' at 1 hert.

![Soil Schematic](soil.png)

## Example

### `air`

```
./air

{"Name": "Air", "Timestamp": 0, "Temperature": 67.4, "Humidity": 0.58}
```

### `soil`

```
./soil

{"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
```
