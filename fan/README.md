# `fan`

`fan` connects STDOUTs from listed out-programs to STDINs of listed in-programs.

## Documentation

`make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make fan` builds `fan`.
* `make test` builds `fan`'s tests.

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

```
cd build && ./test
```

Test-cases are stored in the 'data' directory.
