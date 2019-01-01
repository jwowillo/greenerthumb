# `message`

`message` is where `greenerthumb` messages from the ICD are defined and the
bytes-JSON conversion is implemented.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make message` builds `bytes` and `json`.
* `make bytes` builds `bytes`.
* `make json` builds `json`.
* `make test` builds `message`'s tests.

## Running

```
./bytes
```

```
./json
```

Examples are:

```
./bytes

< {"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
0100000000000000003ebd70a410
```

```
./json

< 0100000000000000003ebd70a410
{"Name": "Soil", "Timestamp": 0, "Moisture": 0.37}
```

The examples show the bytes written and received as hex-strings for
documentation purposes.

## Testing

```
cd build && ./test
```
