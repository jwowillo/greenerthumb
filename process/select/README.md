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

< {"Header": {"Name": "A"}, "1": 1}
{"Header": {"Name": "A"}, "1": 1}
< {"Header": {"Name": "B"}, "1": 1}
{"Header": {"Name": "B"}, "1": 1}
< {"Header": {"Name": "C"}, "1": 1}
```

## Testing

```
cd build && ./test
```
