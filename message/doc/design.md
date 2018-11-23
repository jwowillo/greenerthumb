# `message` Design

`message` is where `greenerthumb` messages from the ICD are defined and the
bytes-JSON conversion is implemented.

## Programs

### `bytes`

`bytes` converts JSON mesages from STDIN to bytes written to STDOUT.

CRC and message errors will be written to STDERR.

### `json`

`json` converts bytes messages from STDIN to JSON written to STDOUT.

Message errors will be written to STDERR.

## Examples

### `bytes`

```
./bytes

< {"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
0100000000000000003ebd70a410
```

### `json`

```
./json

< 0100000000000000003ebd70a410
{"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
```
