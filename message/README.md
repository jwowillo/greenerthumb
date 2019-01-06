# `message`

`message` is where `greenerthumb` messages from the ICD are defined, the
bytes-JSON conversion is implemented, and header wrapping is provided.

## Documentation

* `make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make message` builds `header`, `bytes`, and `json`.
* `make header` builds `header`.
* `make bytes` builds `bytes`.
* `make json` builds `json`.
* `make test` builds `message`'s tests.

## Running

```
./header <name> <sender>
```

The sender can't be longer than 255 characters.

```
./bytes
```

```
./json
```

Examples are:

```
./header name sender

< {"Key": "Value"}
{"Header": {"Name": "name", "Sender": "sender", "Timestamp": 0}, "Key": "Value"}
```

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
