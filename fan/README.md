# `fan`

`fan` connects STDOUTs from listed out-programs to STDINs of listed in-programs.

## Documentation

`make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make fan` builds `fan`.
* `make test` builds `fan`'s tests.
* `make doc` runs `make` in the 'doc' directory.

## Running

```
./fan [--out '<out>'...] [--in '<in>'...]
```

An example is:

```
./fan --out 'echo a' --out 'echo b' --in 'cat' --in 'cat'

a
a
b
b
```

## Testing

All test-cases from the 'data' directory can be run with:

```
cd build && ./test
```
