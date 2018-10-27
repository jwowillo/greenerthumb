# `message` Design

`message` is where `greenerthumb` messages from the ICD are defined and the
bytes-JSON conversion is impmlemented.

## Program

### `bytes`

`bytes` converts JSON mesages from STDIN to bytes written to STDOUT.

CRC and message errors will be written to STDERR.

### `json`

`json` converts bytes messages from STDIN to JSON written to STDOUT.

Message errors will be written to STDERR.

## Example

### `bytes`

```
./bytes

< {"name": "Soil", "timestamp": 0, "Moisture": 0.37}
010000000000000000a470bd3e10
```

### `json`

```
./json

< 010000000000000000a470bd3e10
{"name": "Soil", "timestamp": 0, "Moisture": 0.37}
```
