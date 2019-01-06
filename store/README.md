# `store`

`store` messages with unique names.

`read` will read messages from a store and `write` will write unique messages to
a store updating repeats.

## Documentation

`make` in the 'doc' directory generates PDF documentation.

## Building

* `make` builds all targets.
* `make store` builds `store`.
* `make read` builds `read`.
* `make write` builds `write`.
* `make test` builds `store`'s tests.

## Running

```
./read <path>
```

```
./store <path>
```

An example is:

```
./write store.db

< {"Header": {"Name": "A"}, "Value": 0}
< {"Header": {"Name": "A"}, "Value": 1}
< {"Header": {"Name": "A"}, "Value": 2}
```

```
./read store.db

{"Header": {"Name": "A"}, "Value": 2}
```

## Testing

```
cd build && ./test
```

Test-cases are stored in the 'data' directory.
