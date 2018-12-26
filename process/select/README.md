# `select`

`select` messages from STDIN with names in an included set.

## Building

* `make` builds all targets.
* `make select` builds `select`.
* `make test` builds `select`'s tests.

## Running

```
./select [--include <include>...]
```

An example is:

```
./select --include "A" --include "B"

< {"Name": "A", "Timestamp": 0, "1": 1}
{"Name": "A", "Timestamp": 0, "1": 1}
< {"Name": "B", "Timestamp": 0, "1": 1}
{"Name": "B", "Timestamp": 0, "1": 1}
< {"Name": "C", "Timestamp": 0, "1": 1}
```

## Testing

```
cd build && ./test
```
